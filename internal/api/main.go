package Api

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/database/dbaccount"
	"elton-okawa/battleship/internal/entity/jwttoken"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/controller/controlaccount"
	"elton-okawa/battleship/internal/usecase/game"
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

var ErrMissingAuthorizationHeader = errors.New("missing 'Authorization' header")
var ErrAuthHeaderNotBearer = errors.New("authorization header does not starts with 'Bearer'")

var skipAuthPathPatterns = map[string][]string{
	"POST": {
		"^/accounts$",
		"^/accounts/actions/login$",
	},
	"GET": {},
	"PUT": {},
}

type BattleshipImpl struct {
	accounts controlaccount.Controller
	games    controller.GamesController
}

func SetupHandler() *echo.Echo {
	accountDao := dbaccount.New("./db/accounts.json")
	gameDao := database.NewGameDao("./db/games.json")

	app := BattleshipImpl{
		accounts: controlaccount.New(ucaccount.New(accountDao)),
		games:    controller.NewGamesController(game.NewGameUseCase(gameDao)),
	}

	swagger, err := GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	e := echo.New()
	// Log all requests
	e.Use(echoMiddleware.Logger())

	// The default behavior of AuthenticationValidator is always fail
	// View https://github.com/deepmap/oapi-codegen/issues/221
	validatorOptions := &middleware.Options{}
	validatorOptions.Options.AuthenticationFunc = openapi3filter.NoopAuthenticationFunc

	// Use oapi-generator validation middleware to check all requests
	// against the OpenAPI schema.
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions))

	config := echoMiddleware.JWTConfig{
		// TODO ideally we should use api.yaml to know which endpoints do not need auth
		Skipper: func(c echo.Context) bool {
			patterns := skipAuthPathPatterns[c.Request().Method]
			path := c.Path()
			for _, pattern := range patterns {
				if match, err := regexp.MatchString(pattern, path); err == nil {
					if match {
						return true
					}
				} else {
					panic(err)
				}
			}

			return false
		},
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			claim, err := jwttoken.Validate(auth)

			// Resulting claim will be available in handler's context "user" key
			return claim, err
		},
	}
	e.Use(echoMiddleware.JWTWithConfig(config))

	RegisterHandlers(e, &app)

	return e
}

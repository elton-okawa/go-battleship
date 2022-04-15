package Api

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/database/dbaccount"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/interface_adapter/controller/controlaccount"
	"elton-okawa/battleship/internal/usecase/game"
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

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
	// Use oapi-generator validation middleware to check all requests
	// against the OpenAPI schema.
	e.Use(middleware.OapiRequestValidator(swagger))

	RegisterHandlers(e, &app)

	return e
}

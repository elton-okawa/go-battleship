package Api

import (
	"elton-okawa/battleship/internal/database"
	"elton-okawa/battleship/internal/interface_adapter/controller"
	"elton-okawa/battleship/internal/use_case/account"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type BattleshipImpl struct {
	accounts controller.AccountController
}

func SetupHandler() *echo.Echo {
	accountDao := database.NewAccountDao("./db/accounts.json")

	app := BattleshipImpl{
		accounts: controller.NewAccountController(account.NewAccountUseCase(accountDao)),
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

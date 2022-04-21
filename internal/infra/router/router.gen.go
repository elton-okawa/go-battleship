// Package router provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package router

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new account
	// (POST /accounts)
	CreateAccount(ctx echo.Context) error
	// Perform authentication and receive a jwt token
	// (POST /accounts/actions/login)
	AccountLogin(ctx echo.Context) error
	// Start a new game
	// (POST /games)
	CreateGame(ctx echo.Context) error
	// Shoot
	// (POST /games/{id}/actions/shoot)
	GameShoot(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAccount(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateAccount(ctx)
	return err
}

// AccountLogin converts echo context to params.
func (w *ServerInterfaceWrapper) AccountLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountLogin(ctx)
	return err
}

// CreateGame converts echo context to params.
func (w *ServerInterfaceWrapper) CreateGame(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateGame(ctx)
	return err
}

// GameShoot converts echo context to params.
func (w *ServerInterfaceWrapper) GameShoot(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GameShoot(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/accounts", wrapper.CreateAccount)
	router.POST(baseURL+"/accounts/actions/login", wrapper.AccountLogin)
	router.POST(baseURL+"/games", wrapper.CreateGame)
	router.POST(baseURL+"/games/:id/actions/shoot", wrapper.GameShoot)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xWwW7jNhD9FWLaoxA52xYodEsKNMgiQIN1ix4CH2hqIjMrkVpyGK9h6N+LoWhFsuU0",
	"hxroTaaHw8c3b95wD2vp8VHSBgrIX68hA2Wb1ho05KHYg1cbbGT8fLSebpSywZD/gt8CeuLl1tkWHWmM",
	"QbWttOGPRpsHNBUn/iUD2rUIBXhy2lTQZdBK77fWlUehv56Edhk4/Ba0wxKKp5R/tH817LDrF1TEyadI",
	"fWuNx1OoOh5+gmy4wftAdAmH2HMQ7mSDN4q0NcuNtXSWM2XrRINuQgPFYsinDWGFjhM6u/23oCOEvCOL",
	"yc8BfGD4//9SJpjn6ojfW+3Q38QrlOiV0y2TDgX8ZfR3QbpBoY1oPMwRS/YrfqDefVg2Om0WrrPrGpvP",
	"3ppToCWS1PWs6DxJCn70lwnNOuHTVOMH8MWwIVN2OO0UJR+HKjhNuyV3dw/uFqVDdxO4dHtYx1+/W9dI",
	"ggI+//0np47RUKR/39jcELXQceKNZSVBbZWs43cG2jzb08osKZQ7GG4Hd1bcSqIa/Ua3kMErOt9HXl8t",
	"rhZMhG3RyFZDAVE0tIm4c5kaPfJteyEz65JPui+hgN8cSsJkCNDThp5ubbnr288QmrhPtm2tVdyZv6Qa",
	"9g7IXz86fIYCfsjfLDJP/pjPmWM3rRG5gHGhF3IE/GlxfSEIqVsihin1KUaoSEspfFAKvX8Odb1jnn9e",
	"LP47TKN+mIFyb15lrUuhTRumyoTiaZWBD00j3W6ooZDC4FbIoZQkK8/qP6ysOMcgiVxG9/X5YGXzAkmM",
	"PCRDupQ+Jnb7IXEsLnH+eWUsR0oQta0q5NL0mvjptIfvjbLOoSIRPDphnRis/L1SPqJ7tq4RMtAGDaX7",
	"CGlK4VChfuUyv2xJHCx3vsiVTM71XtPzCIb5npvepQ9PwAeoS5KOkuiqPtUBTPw5QpLvddkNgvM88M+D",
	"Y1jxTRCdzMkGCR2nPYb1Bb0NTqGIzw3WcHQ+yMDIaMVxfSqjbCSJ0Ty+Pp3Hq8tp/czb5+Oqn/Lwx9fj",
	"yiT2jsvRHcZUXOXB5HkEjrQ4HXVPK2bBb2XFz4ECPl0t3rIeo7iLkhu47xXYZft5g30LPAi3W3X/BAAA",
	"//8PkMA6cQsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

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

	"H4sIAAAAAAAC/8xWwW7jNhD9FWLaoxDZi21R6JYUaLCLAA3WLXoIfKCpicysRGrJYbyGoX8vhqIVyZbT",
	"HJqiN4oaDh/fPD7OATbS472kLRSQPy8hA2Wb1ho05KE4gFdbbGQc3ltP10rZYMh/wW8BPfF062yLjjTG",
	"oNpW2vCg0eYOTcWJf8qA9i1CAZ6cNhV0GbTS+5115UnoL2ehXQYOvwXtsITiIeUfrV8PK+zmCRVx8ilS",
	"31rj8RyqjpufIRtO8DoQXcIx9hKEW9ngtSJtzWprLV3kTNk60aCb0ECxGPJpQ1ih44TO7v4p6AQhr8hi",
	"8ksA7xj+/7+UCealOuL3Vjv01/EIJXrldMukQwF/Gv1dkG5QaCMaDxk8WtdI6ln7+SPMMU32K75BAH1Y",
	"Ntp+Fr+zmxqbz96ac+QlbkI1K8ISSep69pcnScGPfo2ha6rxDdBj2JBq2O78ALwfquA07VfsBD3uG5QO",
	"3XXgMh9gE79+OxL7+a8/OHWMhiL9fSF6S9RCx4m3llUHtVWyjuMMtHm051VcUSj3MJwObq24kUQ1+q1u",
	"IYNndL6PXF4trhZMhG3RyFZDAVFgtI24c5lMIZbC9qLngkje6VMJBfzqUBIm84CeNvR0Y8t9f1UNoYnr",
	"ZNvWWsWV+VMqb++WPPrR4SMU8EP+Yqd58tJ8zki7aY3IBYwTvegj4A+L5TtBSDcrYphSn2KEirSUwgel",
	"0PvHUNd75vnjYvHvYRpdlRkon8yzrHUptGnDVJlQPKwz8KFppNsPNRRSGNwJOZSSZOVZ/ceZNecYJJHL",
	"6NQ+H2xvXiCJkbtkXu+lj4k1v0kci/fY/7IyViMliNpWFXJpek0s/ztNKOscKhLBoxPWieExeU0g9+j4",
	"HRAy0BYNJWhCmlI4VKifWTxPOxJHj5+XTiWTH75mJdwEwPxNnp6lD0/AB6grko6SlKs+1RFM/BwhyQ+6",
	"7AYZe245LoNjWLErif7oZIOEjtOewvqC3ganUMSGh29G9FPIwMho8HF+Ks5sVN1RR7A87wjW73eDLnRf",
	"b79LUx5+/3pamcTeaTm64+MXZ/m58/ywjrQ4fUAf1syC38mKH/ECPlwtXrKeoriNkhu47xXYZYd5234J",
	"PAq3W3d/BwAA//8uGHNS8wsAAA==",
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

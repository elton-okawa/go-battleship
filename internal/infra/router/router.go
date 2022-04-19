package router

//go:generate oapi-codegen --config ./router.gen-config.yaml ../../../api/api.yaml

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

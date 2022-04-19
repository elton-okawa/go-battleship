package api

//go:generate oapi-codegen --config ./api.gen-config.yaml ../../api/api.yaml

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

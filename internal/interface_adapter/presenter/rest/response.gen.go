// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package rest

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// PostAccountsRequest defines model for PostAccountsRequest.
type PostAccountsRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostAccountsResponse defines model for PostAccountsResponse.
type PostAccountsResponse struct {
	Id    string `json:"id"`
	Login string `json:"login"`
}

// PostGameActionShootRequest defines model for PostGameActionShootRequest.
type PostGameActionShootRequest struct {
	Col int `json:"col"`
	Row int `json:"row"`
}

// PostLoginRequest defines model for PostLoginRequest.
type PostLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostLoginResponse defines model for PostLoginResponse.
type PostLoginResponse struct {
	// Unix time in ms
	ExpiresAt int64  `json:"expiresAt"`
	Token     string `json:"token"`
}

// ProblemJson defines model for ProblemJson.
type ProblemJson struct {
	Debug  *string `json:"debug,omitempty"`
	Detail string  `json:"detail"`
	Status int     `json:"status"`
	Title  string  `json:"title"`
}

// CreateAccountJSONBody defines parameters for CreateAccount.
type CreateAccountJSONBody PostAccountsRequest

// AccountLoginJSONBody defines parameters for AccountLogin.
type AccountLoginJSONBody PostLoginRequest

// GameShootJSONBody defines parameters for GameShoot.
type GameShootJSONBody PostGameActionShootRequest

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody CreateAccountJSONBody

// AccountLoginJSONRequestBody defines body for AccountLogin for application/json ContentType.
type AccountLoginJSONRequestBody AccountLoginJSONBody

// GameShootJSONRequestBody defines body for GameShoot for application/json ContentType.
type GameShootJSONRequestBody GameShootJSONBody

// Package controlaccount provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package controlaccount

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

// PostLoginRequest defines model for PostLoginRequest.
type PostLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// PostLoginResponse defines model for PostLoginResponse.
type PostLoginResponse struct {
	// Unix time in ms
	ExpiresAt int    `json:"expiresAt"`
	Token     string `json:"token"`
}

// CreateAccountJSONBody defines parameters for CreateAccount.
type CreateAccountJSONBody PostAccountsRequest

// AccountLoginJSONBody defines parameters for AccountLogin.
type AccountLoginJSONBody PostLoginRequest

// CreateAccountJSONRequestBody defines body for CreateAccount for application/json ContentType.
type CreateAccountJSONRequestBody CreateAccountJSONBody

// AccountLoginJSONRequestBody defines body for AccountLogin for application/json ContentType.
type AccountLoginJSONRequestBody AccountLoginJSONBody

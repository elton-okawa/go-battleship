package account

import (
	"context"
	"elton-okawa/battleship/internal/e2e"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

type TestSuite struct {
	suite.Suite
	clt e2e.ClientWithResponsesInterface
	svr *httptest.Server
	req e2e.CreateAccountJSONRequestBody
}

func (s *TestSuite) SetupSuite() {
	s.svr = e2e.SetupTestServer()

	s.req = e2e.CreateAccountJSONRequestBody{
		Login:    "username",
		Password: "password",
	}

	clt, _ := e2e.NewClientWithResponses(s.svr.URL)
	s.clt = clt
}

func (s *TestSuite) SetupTest() {
	e2e.CleanupDatabase()
}

func (s *TestSuite) TearDownSuite() {
	s.svr.Close()
}

func (s TestSuite) TestPostAccount() {
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(res.StatusCode(), 201)

	s.NotEmpty(res.JSON201.Id)
	s.Equal("username", res.JSON201.Login)
}

func (s TestSuite) TestPostAccount_LoginFewerThanFiveChar() {
	s.req.Login = "user"
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(res.StatusCode(), 400)

	s.Contains(res.JSON400.Detail, "login")
}

func (s TestSuite) TestPostAccount_PasswordFewerThanEightChar() {
	s.req.Password = "pass"
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(res.StatusCode(), 400)

	s.Contains(res.JSON400.Detail, "password")
}

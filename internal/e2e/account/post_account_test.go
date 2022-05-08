package account

import (
	"context"
	"elton-okawa/battleship/internal/e2e"
	"elton-okawa/battleship/internal/e2e/setup"
	"elton-okawa/battleship/internal/test"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite_PostAccount(t *testing.T) {
	suite.Run(t, new(TestPostAccountSuite))
}

type TestPostAccountSuite struct {
	suite.Suite
	clt e2e.ClientWithResponsesInterface
	svr *httptest.Server
	req e2e.CreateAccountJSONRequestBody
}

func (s *TestPostAccountSuite) SetupSuite() {
	svr, _ := setup.TestServer()
	s.svr = svr

	s.req = e2e.CreateAccountJSONRequestBody{
		Login:    "username",
		Password: "password",
	}

	clt, _ := e2e.NewClientWithResponses(s.svr.URL)
	s.clt = clt
}

func (s *TestPostAccountSuite) SetupTest() {
	test.CleanupDatabase()
}

func (s *TestPostAccountSuite) TearDownSuite() {
	s.svr.Close()
}

func (s TestPostAccountSuite) TestPostAccount() {
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(201, res.StatusCode())

	s.NotEmpty(res.JSON201.Id)
	s.Equal("username", res.JSON201.Login)
}

func (s TestPostAccountSuite) TestPostAccount_LoginFewerThanFiveChar() {
	s.req.Login = "user"
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(400, res.StatusCode())

	s.Contains(res.JSON400.Detail, "login")
}

func (s TestPostAccountSuite) TestPostAccount_PasswordFewerThanEightChar() {
	s.req.Password = "pass"
	res, err := s.clt.CreateAccountWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)
	s.Equal(400, res.StatusCode())

	s.Contains(res.JSON400.Detail, "password")
}

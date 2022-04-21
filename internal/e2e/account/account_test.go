package account

import (
	"context"
	"elton-okawa/battleship/internal/e2e"
	"elton-okawa/battleship/internal/infra/router"
	"fmt"
	"net/http/httptest"
	"path/filepath"
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

var path, _ = filepath.Abs(filepath.Join("..", "..", "..", "db", "test"))

func (s *TestSuite) SetupSuite() {
	opt := router.Options{
		Db: router.DBOptions{
			Path: path,
		},
	}

	rt := router.Setup(opt)
	svr := httptest.NewServer(rt)
	s.svr = svr

	s.req = e2e.CreateAccountJSONRequestBody{
		Login:    "username",
		Password: "password",
	}

	clt, _ := e2e.NewClientWithResponses(svr.URL)
	s.clt = clt
	fmt.Printf("Test server listening to '%s'\n", svr.URL)
}

func (s *TestSuite) SetupTest() {
	e2e.CleanupDatabase(path)
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

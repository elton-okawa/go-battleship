package account

import (
	"context"
	"elton-okawa/battleship/internal/e2e"
	"elton-okawa/battleship/internal/e2e/setup"
	"elton-okawa/battleship/internal/entity/account"
	"elton-okawa/battleship/internal/usecase/ucaccount"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestSuite_Login(t *testing.T) {
	suite.Run(t, new(TestLoginSuite))
}

type TestLoginSuite struct {
	suite.Suite
	clt        e2e.ClientWithResponsesInterface
	svr        *httptest.Server
	accountDao ucaccount.Repository
	req        e2e.AccountLoginJSONRequestBody
}

func (s *TestLoginSuite) SetupSuite() {
	svr, db := setup.TestServer()

	s.svr = svr
	s.accountDao = db.Account
	s.req = e2e.AccountLoginJSONRequestBody{
		Login:    "username",
		Password: "password",
	}

	clt, _ := e2e.NewClientWithResponses(s.svr.URL)
	s.clt = clt
}

func (s *TestLoginSuite) SetupTest() {
	setup.CleanupDatabase()

	acc, _ := account.New(s.req.Login, s.req.Password)
	s.accountDao.Save(acc)
}

func (s *TestLoginSuite) TearDownSuite() {
	s.svr.Close()
}

func (s TestLoginSuite) TestLogin() {
	res, err := s.clt.AccountLoginWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)

	s.Equal(200, res.StatusCode())
	s.NotEmpty(res.JSON200.Token)
	s.Less(time.Now().Unix(), res.JSON200.ExpiresAt)
}

func (s TestLoginSuite) TestLogin_IncorrectLogin() {
	s.req.Login = "others"
	res, err := s.clt.AccountLoginWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)

	s.Equal(401, res.StatusCode())
}

func (s TestLoginSuite) TestLogin_IncorrectPassword() {
	s.req.Password = "wrong-password"
	res, err := s.clt.AccountLoginWithResponse(context.TODO(), s.req)
	s.Nilf(err, "unexpected error %v", err)

	s.Equal(401, res.StatusCode())
}

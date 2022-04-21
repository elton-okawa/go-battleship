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
}

var path, _ = filepath.Abs(filepath.Join("..", "..", "..", "db", "test"))

func (ts *TestSuite) SetupSuite() {
	opt := router.Options{
		Db: router.DBOptions{
			Path: path,
		},
	}

	rt := router.Setup(opt)
	svr := httptest.NewServer(rt)
	ts.svr = svr

	clt, _ := e2e.NewClientWithResponses(svr.URL)
	ts.clt = clt
	fmt.Printf("Test server listening to '%s'\n", svr.URL)
}

func (ts *TestSuite) SetupTest() {
	e2e.CleanupDatabase(path)
}

func (ts *TestSuite) TearDownSuite() {
	ts.svr.Close()
}

func (ts *TestSuite) TestPostAccount() {
	res, err := ts.clt.CreateAccountWithResponse(context.TODO(), e2e.CreateAccountJSONRequestBody{
		Login:    "username",
		Password: "password",
	})
	ts.Nilf(err, "unexpected error %v", err)
	ts.Equal(res.StatusCode(), 201)

	ts.NotEmpty(res.JSON201.Id)
	ts.Equal("username", res.JSON201.Login)
}

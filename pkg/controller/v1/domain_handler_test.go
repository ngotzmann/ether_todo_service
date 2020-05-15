package v1

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/ngotzmann/gorror"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testLoginSucess(e *httpexpect.Expect) {
	e.POST("/login").
				WithQuery("username", "jdoe").
					WithQuery("password", "Test1234#").
						Expect().
							Status(http.StatusOK)
}

func testReadSucessfull(e *httpexpect.Expect) {

	//Get jwt token for authentication
	res := e.POST("/login").
		WithQuery("username", "jdoe").
			WithQuery("password", "Test1234#").
				Expect().Status(http.StatusOK).
					JSON().Object()
	res.Keys().ContainsOnly("token")
	token := res.Value("token").String().Raw()

	//Get user with username "jdoe" from restricted path
	e.GET("/r/user/jdoe").WithHeader("Authorization", "Bearer " + token).
			Expect().
				Status(http.StatusOK).
					JSON().Object().
						Values().Contains("jdoe")

}

func testLogoutSucess(e *httpexpect.Expect) {
	res := e.POST("/login").
		WithQuery("username", "jdoe").
		WithQuery("password", "Test1234#").
		Expect().Status(http.StatusOK).
		JSON().Object()
	res.Keys().ContainsOnly("token")
	token := res.Value("token").String().Raw()

	e.GET("/r/logout").WithHeader("Authorization", "Bearer " + token).
		Expect().
			Status(http.StatusOK)
}

func testLoginUnsucessfull(e *httpexpect.Expect) {
	e.POST("/login").
			WithQuery("password", "Test1234").
				WithQuery("username", "jdoe").
					Expect().Status(http.StatusUnauthorized)
}

func TestEchoClient(t *testing.T) {
	gorror.Init("../../../config/")
	handler := EchoHandler("../../../config")

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	testLoginSucess(e)
	testLoginUnsucessfull(e)
	testReadSucessfull(e)
	testLogoutSucess(e)
}
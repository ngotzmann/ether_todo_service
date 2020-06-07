package v1

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/ngotzmann/gorror"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TODO: Test deletion of tasks

//TODO: Create Test
// postiv test, create new entry, read it, delete it
// positiv test, create new entry, read, recreate the same entry(check time), read, delete
// negativ test, validation failed


//TODO: implement negativ test
func testFindListByNameSuccessful(e *httpexpect.Expect) {
	e.GET("/todo/list/ketchup").
		Expect().
			Status(http.StatusOK).
				JSON().Object().
					Values().Contains("ketchup")
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

	testFindListByNameSuccessful(e)
}
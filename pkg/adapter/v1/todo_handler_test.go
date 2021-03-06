package v1

import (
	"errors"
	"ether_todo/pkg/glue"
	"ether_todo/pkg/glue/config"
	"ether_todo/pkg/injector"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	"github.com/kataras/i18n"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Create list -> check if list is created -> delete list
func positivIntegrationTestCreate(e *httpexpect.Expect) {
	name := "int_positiv_integration_test_create"
	testCreateListSuccessful(e, &reqList{
		Name: name,
		LiveTime: "keep",
	})
	testFindListByNameSuccessful(e, name)
	err := s.DeleteListByName(name)
	if err != nil {
		log.Error(err)
	}
}
//Create list -> check if list is created -> add tasks -> check if task is created -> create list with same name -> check if no tasks exists
//-> delete listCleanOutatedLists
func positivIntegrationTestOverwrite(e *httpexpect.Expect) {
	name := "int_positiv_integration_test_overwrite"
	testCreateListSuccessful(e, &reqList{
		Name: name,
		LiveTime: "keep",
	})
	testFindListByNameSuccessful(e, name)
	//TODO: add tasks
	testCreateListSuccessful(e, &reqList{
		Name: name,
		LiveTime: "keep",
	})
	testFindListByNameSuccessful(e, name)

	err := s.DeleteListByName(name)
	if err != nil {
		log.Error(err)
	}
}

func negativIntegrationTestCreateValidationFailed(e *httpexpect.Expect) {
	req := reqList{
		Name:     "",
		LiveTime: "keep",
	}
	/*res := e.POST("/todo/list").
		WithForm(req).Expect().Status(http.StatusBadRequest)*/
	path := "/todo/list"
	res := e.POST(path).WithForm(req).Expect().Status(400) //.Status(http.StatusBadRequest)

	fmt.Println(res)
}

func testCreateListSuccessful(e *httpexpect.Expect, reqBody *reqList) {
	 e.POST("/todo/list").
	 	WithForm(reqBody).
	Expect().Status(http.StatusOK)
}

func testFindListByNameSuccessful(e *httpexpect.Expect, name string) {
	res := e.GET("/todo/list/" + name).
		Expect().
			Status(http.StatusOK).
				JSON().Object()

	//check if created_at date is set.
	res.Value("created_at").NotNull()

	//ckeck if id is a valid uuid, so it was set in the db
	id := res.Value("id").String().Raw()
	_, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
	}
}

func TestEchoClient(t *testing.T) {
	_, err := i18n.New(i18n.Glob("./locales/*/*"), "en-US")
	if err != nil {
		panic(err)
	}
	err = errors.New(i18n.Tr("en-US","ValidationError") + " " + "errMsgs")

	config.CustomCfgLocation = "../../../config/local"
	h := glue.DefaultHttpServer()
	h = Endpoints(h)
	injector.TodoService().Migration()
	srv := httptest.NewServer(h)
	defer srv.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  srv.URL,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
	positivIntegrationTestCreate(e)
	positivIntegrationTestOverwrite(e)
	negativIntegrationTestCreateValidationFailed(e)
}

type reqList struct {
	Name string `json:"name"`
	LiveTime string `json:"liveTime"`
}

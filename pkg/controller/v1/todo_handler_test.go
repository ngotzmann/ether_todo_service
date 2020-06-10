package v1

import (
	"ether_todo/pkg/controller/persistence"
	"ether_todo/pkg/todo"
	"github.com/gavv/httpexpect/v2"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/ngotzmann/gommon"
	"github.com/ngotzmann/gormmon"
	"github.com/ngotzmann/gorror"
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
	repo := persistence.NewTodoListRepo()
	uc := todo.NewUsecase(repo, todo.NewService(repo))
	err := uc.DeleteListByName(name)
	if err != nil {
		log.Error(err)
	}
}
//Create list -> check if list is created -> add tasks -> check if task is created -> create list with same name -> check if no tasks exists
//-> delete list
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

	repo := persistence.NewTodoListRepo()
	uc := todo.NewUsecase(repo, todo.NewService(repo))
	err := uc.DeleteListByName(name)
	if err != nil {
		log.Error(err)
	}
}

func negativIntegrationTestCreateValidationFailed(e *httpexpect.Expect) {
	e.POST("/todo/list").
		WithForm(&reqList{
		Name:     "",
		LiveTime: "keep",
	}).
		Expect().Status(http.StatusConflict)
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
	cP := "../../../config/"
	gorror.Init(cP)
	c := gommon.NewConfig(cP)
	gormmon.InitGormDB(gormmon.GormConfig{
		Host:               c.Database.Address,
		Port:               c.Database.Port,
		DBName:             c.Database.Database,
		Username:           c.Database.User,
		Password:           c.Database.Password,
		MaxIdleConnections: c.Database.MaxIdleConnections,
		ShouldLog:          c.Database.Logging,
	})
	h := EchoHandler(cP)
	
	server := httptest.NewServer(h)
	defer server.Close()

	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  server.URL,
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
	LiveTime string `json:"live_time"`
}

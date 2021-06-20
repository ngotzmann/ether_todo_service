package todo

import (
	"github.com/kataras/i18n"
	"testing"
	"time"

	"github.com/google/uuid"
)

func getList() *List {
	return &List{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
		Name: 	   "TestList",
		LiveTime:  "keep",
	}
}

func before()  {
	_, err := i18n.New(i18n.Glob("./locales/*/*"), "en-US")
	if err != nil {
		panic(err)
	}

}

func TestModelValidation_ExpectSuccess(t *testing.T) {
	before()
	l := getList()
	err := l.Validation()
	if err != nil {
		t.Log(err)
		t.Errorf("User validation was incorrect!")
	}
}

func TestModelValidation_ExpectFailed(t *testing.T) {
	l := getList()
	l.Name = ""
	err := l.Validation()
	if err != nil {
		t.Log("Test Success")
		t.Log(err)
	} else {
		t.Errorf("User validation was incorrect!")
	}
}

package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ngotzmann/gorror"
)

func getTestModel() *Model {
	return &Model{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

func TestModelValidationSuccess(t *testing.T) {
	gorror.Init("")
	u := getTestModel()
	err := u.Validation()
	if err != nil {
		t.Log(err)
		t.Errorf("User validation was incorrect!")
	}
}

func TestModelValidationFailed(t *testing.T) {
	gorror.Init("")
	u := getTestModel()
	u.ID = uuid.UUID{}
	err := u.Validation()
	if err != nil {
		t.Log("Test Success")
		t.Log(err)
	} else {
		t.Errorf("User validation was incorrect!")
	}
}

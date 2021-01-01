package todo

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func getTestModel() *List {
	return &List{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
}

func TestModelValidationSuccess(t *testing.T) {
	u := getTestModel()
	err := u.Validation()
	if err != nil {
		t.Log(err)
		t.Errorf("User validation was incorrect!")
	}
}

func TestModelValidationFailed(t *testing.T) {
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

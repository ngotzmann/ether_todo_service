package todo

import (
	"ether_todo/pkg/controller/persistence/gormmon"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ngotzmann/gorror"
)

type Task struct {
	gormmon.Base
	Task       string    `json:"task" validate:"required"`
	ChildTasks []Task    `json:"child_tasks"`
	DueDate    time.Time `json:"due_date" validate:"gte"`
	Color      string    `json:"color" validate:"required,hexcolor"`
	Done 	   bool      `json:"done"`
	ListId    uuid.UUID  `gorm:"list_id"`
}

func (m *Task) Validation() error {
	v := validator.New()
	var errMsgs string
	err := v.Struct(m)
	if err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			specErrMsg := fmt.Sprintf("%v", fe)
			errMsgs += specErrMsg + "; "
		}
		err = gorror.CreateError(gorror.ValidationError, errMsgs)
	}
	return err
}

package todo

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ngotzmann/gorror"
)

type Task struct {
	ID         uuid.UUID `json:"id"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
	Task       string    `json:"task" validate:"required"`
	ChildTasks []Task    `json:"child_tasks"`
	DueDate    time.Time `json:"due_date" validate:"gte"`
	Color      string    `json:"color" validate:"required,hexcolor"`
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

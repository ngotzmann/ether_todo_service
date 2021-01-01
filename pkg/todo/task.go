package todo

import (
	"errors"
	"fmt"
	"github.com/kataras/i18n"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`

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
		err = errors.New(i18n.Tr("ValidationError", "en-US") + " " + errMsgs)
	}
	return err
}

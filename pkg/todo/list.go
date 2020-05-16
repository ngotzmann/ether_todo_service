package todo

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ngotzmann/gorror"
)

type LiveTime string

const (
	Day   LiveTime = "tmp"
	Month LiveTime = "mth"
	Year  LiveTime = "keep"
)

type List struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name" validate:"required"`
	Tasks     []Task    `json:"tasks"`
	LiveTime  LiveTime  `json:"live_time, validate:"required""`
}

func (m *Model) Validation() error {
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

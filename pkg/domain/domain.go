package domain

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/ngotzmann/gorror"
	"time"
)

type Model struct {
	ID uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
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
package todo

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kataras/i18n"
	"time"
)

type LiveTime string

const (
	Day   LiveTime = "tmp"
	Month LiveTime = "mth"
	Year  LiveTime = "keep"
)

type List struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`

	Name      string    `json:"name" validate:"required" gorm:"name"`
	Tasks     []Task    `json:"tasks" gorm:"foreignkey:list_id"`
	LiveTime  LiveTime  `json:"liveTime" validate:"required"`
}

func (l *List) Validation() error {
	err :=l.ValidateLiveTimeEnum()
	if err != nil {
		return err
	}
	v := validator.New()
	var errMsgs string
	err = v.Struct(l)
	if err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			specErrMsg := fmt.Sprintf("%v", fe)
			errMsgs += specErrMsg + "; "
		}
		err = errors.New(i18n.Tr("en-US", "ValidationError") + " " + errMsgs)
	}
	return err
}

func (l *List) ValidateLiveTimeEnum() error {
	switch l.LiveTime {
	case Day, Month, Year:
		return nil
	}
	return errors.New(i18n.Tr("en-US","ValidationError") + " not allowed LiveTime value")
}
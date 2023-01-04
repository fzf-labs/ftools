package custom

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type DateLt struct {
}

func NewDateLt() *DateLt {
	return &DateLt{}
}

func (d *DateLt) Tag() string {
	return "DateLt"
}

func (d *DateLt) ZhTranslation() string {
	return "{0} 必须要小于当前日期"
}

func (d *DateLt) EnTranslation() string {
	return "{0} must be less than the current date"
}

func (d *DateLt) Validate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.After(time.Now()) {
		return false
	}
	return true
}

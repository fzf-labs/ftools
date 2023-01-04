package custom

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type DateGt struct {
}

func NewDateGt() *DateGt {
	return &DateGt{}
}

func (d *DateGt) Tag() string {
	return "DateGt"
}

func (d *DateGt) ZhTranslation() string {
	return "{0} 必须要大于当前日期"
}

func (d *DateGt) EnTranslation() string {
	return "{0} must be greater than the current date"

}

func (d *DateGt) Validate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	if date.Before(time.Now()) {
		return false
	}
	return true
}

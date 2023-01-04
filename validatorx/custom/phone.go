package custom

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Phone struct {
}

func NewPhone() *Phone {
	return &Phone{}
}

func (p *Phone) Tag() string {
	return "phone"
}

func (p *Phone) ZhTranslation() string {
	return "{0} 错误的手机格式"

}

func (p *Phone) EnTranslation() string {
	return "{0} wrong mobile phone format"
}

func (p *Phone) Validate(fl validator.FieldLevel) bool {
	return IsPhone(fl.Field().String())
}

func IsPhone(phone string) bool {
	regular := `^13[\d]{9}$|^14[5,7]{1}\d{8}$|^15[^4]{1}\d{8}$|^16[\d]{9}$|^17[0,2,3,5,6,7,8]{1}\d{8}$|^18[\d]{9}$|^19[\d]{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// IsPhoneLoose 宽松的手机号验证
// 13, 14, 15, 16, 17, 18, 19 can pass the verification (只要满足 13、14、15、16、17、18、19开头的11位数字都可以通过验证)
func IsPhoneLoose(phone string) bool {
	regular := `^1(3|4|5|6|7|8|9)\d{9}$`

	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

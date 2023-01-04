package validatorx

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Validation interface {
	Tag() string                           //标签
	ZhTranslation() string                 //中文翻译
	EnTranslation() string                 //英语翻译
	Validate(fl validator.FieldLevel) bool //验证函数
}

func (v *ValidatorX) RegisterValidate(validations ...Validation) error {
	for _, validation := range validations {
		err := v.RegisterValidation(validation)
		if err != nil {
			return err
		}
		err = v.RegisterTranslation(validation)
		if err != nil {
			return err
		}
	}
	return nil
}
func (v *ValidatorX) RegisterValidation(validation Validation) error {
	err := v.Validator.RegisterValidation(validation.Tag(), validation.Validate)
	if err != nil {
		return err
	}
	return nil
}

// RegisterTranslation 注册翻译器
func (v *ValidatorX) RegisterTranslation(validation Validation) error {
	if err := v.Validator.RegisterTranslation(validation.Tag(), v.Trans["zh"], registerTranslator(validation.Tag(), validation.ZhTranslation()), translate); err != nil {
		return err
	}
	if err := v.Validator.RegisterTranslation(validation.Tag(), v.Trans["en"], registerTranslator(validation.Tag(), validation.EnTranslation()), translate); err != nil {
		return err
	}
	return nil
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

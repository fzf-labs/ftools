package validatorx

import (
	"fmt"
	"ftools/validatorx/custom"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

var Validator = NewValidator()

type ValidatorX struct {
	Validator *validator.Validate      //验证器
	Uni       *ut.UniversalTranslator  //通用翻译器
	Trans     map[string]ut.Translator //翻译器
}

func NewValidator() *ValidatorX {
	v := ValidatorX{}
	//1.验证器
	v.Validator = validator.New()
	// 注册一个获取json tag的自定义方法
	v.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	//2.通用翻译器
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器
	// 第一个参数是备用（fallback）的语言环境,后面的参数是应该支持的语言环境（支持多个）,uni := ut.New(zhT, zhT) 也是可以的
	v.Uni = ut.New(zhT, enT, zhT)
	//3.翻译器
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans
	//注册翻译器
	err := zhTranslations.RegisterDefaultTranslations(v.Validator, v.Trans["zh"])
	if err != nil {
		panic("Validator RegisterDefaultTranslations(zh) failed")
	}
	err = enTranslations.RegisterDefaultTranslations(v.Validator, v.Trans["en"])
	if err != nil {
		panic("Validator RegisterDefaultTranslations(en) failed")
	}
	err = v.RegisterValidate(custom.NewDateGt(), custom.NewDateLt(), custom.NewPhone())
	if err != nil {
		panic(fmt.Sprintf("Validator RegisterValidate failed,err: %s", err))
	}
	return &v
}

// Validate validate
func (v *ValidatorX) Validate(obj interface{}, lang string) error {
	if obj == nil {
		return nil
	}
	lang = acceptLanguage(lang)
	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		err := v.Validator.Struct(value.Elem().Interface())
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				return errors.New(e.Translate(v.Trans[lang]))
			}
			return err
		}
	case reflect.Struct:
		err := v.Validator.Struct(obj)
		if err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				return errors.New(e.Translate(v.Trans[lang]))
			}
			return err
		}
	case reflect.Slice, reflect.Array:
		count := value.Len()
		for i := 0; i < count; i++ {
			if err := v.Validator.Struct(value.Index(i).Interface()); err != nil {
				for _, e := range err.(validator.ValidationErrors) {
					return errors.New(e.Translate(v.Trans[lang]))
				}
				return err
			}
		}
	default:
		return nil
	}
	return nil
}

func acceptLanguage(lang string) string {
	switch lang {
	case "zh-CN":
		return "zh"
	case "en-US":
		return "en"
	default:
		return "zh"
	}
}

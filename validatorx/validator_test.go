package validatorx

import (
	"fmt"
	"testing"
)

func TestNewValidator(t *testing.T) {
	validator := NewValidator()
	type Man struct {
		Name    string `json:"name" `
		Version string `json:"version" validate:"required"`
		Time    string `json:"time" validate:"DateGt"`
		Info    string `json:"info"`
	}
	man := Man{
		Version: "121",
		Info:    "12312",
	}
	err := validator.Validate(man, "zh")

	fmt.Println(err)
}

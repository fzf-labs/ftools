package conv

import (
	"fmt"
	"testing"
)

func TestBase10To64(t *testing.T) {
	to64, err := Base10To64(29)
	fmt.Println(to64)
	fmt.Println(err)
}

func TestBase64To10(t *testing.T) {
	to64, err := Base64To10("0t")
	fmt.Println(to64)
	fmt.Println(err)
}

package osutil

import (
	"fmt"
	"testing"
)

func TestIsWindows(t *testing.T) {
	fmt.Println(IsWindows())
	fmt.Println(IsLinux())
	fmt.Println(IsMac())
}

package oss

import (
	"fmt"
	"testing"
)

func TestAliConfig_DeCrypt(t *testing.T) {
	aliConfig := AliConfig{
		Salt: "xiaoxie",
	}
	crypt, err := aliConfig.DeCrypt("https://666-meishu.oss-cn-shenzhen.aliyuncs.com/fgzs/pro/avatar/fake_user/20220415/6a80ab76-bcac-11ec-959c-00163e02a739.png?encrypt=1&size=150,150")
	if err != nil {
		return
	}
	fmt.Println(crypt)
}

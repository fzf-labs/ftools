package webhook

import (
	"fmt"
	"testing"
)

func TestFeiShu_SendMsg(t *testing.T) {
	feiShu := NewFeiShu(&FeiShuConfig{
		Url:  "https://open.feishu.cn/open-apis/bot/v2/hook/790d95f6-47f8-45b0-8cfe-7635dec6c1d4",
		Sign: "i25mfl9o1MT2vpUOkCtEZc",
	})
	err := feiShu.SendMsg("测试")
	if err != nil {
		return
	}
	fmt.Println(err)
}

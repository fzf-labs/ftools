package speech

import (
	"fmt"
	"github.com/imroc/req/v3"
)

type SpeechResp struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUid int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func GetWord() (word string, err error) {
	var response SpeechResp
	client := req.C()
	resp, err := client.R().SetResult(&response).Post("https://v1.hitokoto.cn/?c=k")
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", fmt.Errorf("bad response status: %s", resp.Status)
	}
	word = response.Hitokoto
	if response.From != "" {
		word += "  ——" + response.From
	}
	return word, nil
}

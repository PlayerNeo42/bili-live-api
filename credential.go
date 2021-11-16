package api

import (
	"fmt"
	json "github.com/json-iterator/go"
	"os"
)

// Credential 保存用户认证信息
type Credential struct {
	SessData string `json:"SESSDATA,omitempty"`
	BiliJct  string `json:"bili_jct,omitempty"`
	Buvid3   string `json:"buvid3,omitempty"`
}

func (c Credential) Cookie() string {
	return fmt.Sprintf("SESSDATA=%s;bili_jct=%s;buvid3=%s", c.SessData, c.BiliJct, c.Buvid3)
}

func (c *Credential) LoadFromLocalFile() *Credential{
	f, err := os.Open("credential.json")
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		panic(err)
	}
	return c
}

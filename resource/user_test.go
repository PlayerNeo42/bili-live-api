package resource

import (
	"encoding/json"
	"testing"
)

func TestUserInfo(t *testing.T) {
	userInfo, err := UserInfo(2075179777)
	userJson, _ := json.Marshal(userInfo)
	println(string(userJson))
	if err != nil {
		t.Error(err)
	}
	if userInfo.Message == "" {
		t.Errorf("获取用户信息失败, %+v", userInfo)
	}

}

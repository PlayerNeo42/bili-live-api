package resource

import (
	"testing"
)

func TestUserInfo(t *testing.T) {
	userInfo, err := UserInfo(1)
	if err != nil {
		t.Error(err)
	}
	if userInfo.Message == "" {
		t.Errorf("获取用户信息失败, %+v", userInfo)
	}
}

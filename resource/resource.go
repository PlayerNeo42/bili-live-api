package resource

import "github.com/go-resty/resty/v2"

var apiClient = resty.New()

func init() {
	apiClient.SetBaseURL(APIUrl)
}

const (
	WSUrl  = "wss://broadcastlv.chat.bilibili.com/sub"
	APIUrl = "https://api.live.bilibili.com"
)

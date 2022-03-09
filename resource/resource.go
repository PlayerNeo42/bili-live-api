package resource

import "github.com/go-resty/resty/v2"

var apiClient = resty.New()

func init() {
	apiClient.SetBaseURL(APIUrl)
}

const (
	// WSUrl B站直播websocket接入地址
	WSUrl = "wss://broadcastlv.chat.bilibili.com/sub"
	// APIUrl B站API地址
	APIUrl = "https://api.live.bilibili.com"
)

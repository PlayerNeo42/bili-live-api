package resource

import "github.com/go-resty/resty/v2"

const (
	// WSUrl B站直播websocket接入地址
	WSUrl = "wss://broadcastlv.chat.bilibili.com/sub"
	// LiveAPIURL B站直播API地址
	LiveAPIURL = "https://api.live.bilibili.com"
	// APIURL B站API地址
	APIURL         = "https://api.bilibili.com"
	VcAPIURL       = "https://api.vc.bilibili.com"
	UserAgentKey   = "User-Agent"
	UserAgentValue = "curl 7.68.0"
)

var (
	liveAPIClient = resty.New().SetHeader(UserAgentKey, UserAgentValue).SetBaseURL(LiveAPIURL)
	apiClient     = resty.New().SetHeader(UserAgentKey, UserAgentValue).SetBaseURL(APIURL)
	vcApiClient   = resty.New().SetHeader(UserAgentKey, UserAgentValue).SetBaseURL(VcAPIURL)
)

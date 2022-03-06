package dto

type HotRankChangedV2 struct {
	Rank        int    `json:"rank"`
	Trend       int    `json:"trend"`
	Countdown   int    `json:"countdown"`
	Timestamp   int    `json:"timestamp"`
	WebUrl      string `json:"web_url"`
	LiveUrl     string `json:"live_url"`
	BlinkUrl    string `json:"blink_url"`
	LiveLinkUrl string `json:"live_link_url"`
	PcLinkUrl   string `json:"pc_link_url"`
	Icon        string `json:"icon"`
	AreaName    string `json:"area_name"`
	RankDesc    string `json:"rank_desc"`
}

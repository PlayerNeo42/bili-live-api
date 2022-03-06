package dto

type HotRankChanged struct {
	Rank        int    `json:"rank"`
	Trend       int    `json:"trend"`
	Countdown   int    `json:"countdown"`
	Timestamp   int    `json:"timestamp"`
	WebURL      string `json:"web_url"`
	LiveURL     string `json:"live_url"`
	BlinkURL    string `json:"blink_url"`
	LiveLinkURL string `json:"live_link_url"`
	PcLinkURL   string `json:"pc_link_url"`
	Icon        string `json:"icon"`
	AreaName    string `json:"area_name"`
	RankDesc    string `json:"rank_desc"`
}

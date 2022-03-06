package dto

type HotRankSettlementV2 struct {
	Rank      int    `json:"rank"`
	Uname     string `json:"uname"`
	Face      string `json:"face"`
	Timestamp int    `json:"timestamp"`
	Icon      string `json:"icon"`
	AreaName  string `json:"area_name"`
	Url       string `json:"url"`
	CacheKey  string `json:"cache_key"`
	DmMsg     string `json:"dm_msg"`
}

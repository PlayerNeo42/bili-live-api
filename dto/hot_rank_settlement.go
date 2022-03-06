package dto

type HotRankSettlement struct {
	AreaName  string `json:"area_name"`
	CacheKey  string `json:"cache_key"`
	DmMsg     string `json:"dm_msg"`
	DmScore   int    `json:"dmscore"`
	Face      string `json:"face"`
	Icon      string `json:"icon"`
	Rank      int    `json:"rank"`
	Timestamp int    `json:"timestamp"`
	Uname     string `json:"uname"`
	Url       string `json:"url"`
}

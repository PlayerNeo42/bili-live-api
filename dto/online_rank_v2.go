package dto

type OnlineRankV2 struct {
	List []struct {
		Uid        int    `json:"uid"`
		Face       string `json:"face"`
		Score      string `json:"score"`
		Uname      string `json:"uname"`
		Rank       int    `json:"rank"`
		GuardLevel int    `json:"guard_level"`
	} `json:"list"`
	RankType string `json:"rank_type"`
}

package dto

type OnlineRankTop3 struct {
	DmScore int `json:"dmscore"`
	List    []struct {
		Msg  string `json:"msg"`
		Rank int    `json:"rank"`
	} `json:"list"`
}

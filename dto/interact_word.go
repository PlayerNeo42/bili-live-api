package dto

type InteractWord struct {
	Contribution struct {
		Grade int `json:"grade"`
	} `json:"contribution"`
	DmScore     int    `json:"dmscore"`
	FansMedal   Medal  `json:"fans_medal"`
	Identities  []int  `json:"identities"`
	IsSpread    int    `json:"is_spread"`
	MsgType     int    `json:"msg_type"`
	RoomID      int    `json:"roomid"`
	Score       int64  `json:"score"`
	SpreadDesc  string `json:"spread_desc"`
	SpreadInfo  string `json:"spread_info"`
	TailIcon    int    `json:"tail_icon"`
	Timestamp   int    `json:"timestamp"`
	TriggerTime int64  `json:"trigger_time"`
	UID         int    `json:"uid"`
	Uname       string `json:"uname"`
	UnameColor  string `json:"uname_color"`
}

package dto

type NoticeMsg struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Full struct {
		HeadIcon    string `json:"head_icon"`
		TailIcon    string `json:"tail_icon"`
		HeadIconFa  string `json:"head_icon_fa"`
		TailIconFa  string `json:"tail_icon_fa"`
		HeadIconFan int    `json:"head_icon_fan"`
		TailIconFan int    `json:"tail_icon_fan"`
		Background  string `json:"background"`
		Color       string `json:"color"`
		Highlight   string `json:"highlight"`
		Time        int    `json:"time"`
	} `json:"full"`
	Half struct {
		HeadIcon   string `json:"head_icon"`
		TailIcon   string `json:"tail_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Time       int    `json:"time"`
	} `json:"half"`
	Side struct {
		HeadIcon   string `json:"head_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Border     string `json:"border"`
	} `json:"side"`
	RoomID     int    `json:"roomid"`
	RealRoomID int    `json:"real_roomid"`
	MsgCommon  string `json:"msg_common"`
	MsgSelf    string `json:"msg_self"`
	LinkURL    string `json:"link_url"`
	MsgType    int    `json:"msg_type"`
	ShieldUID  int    `json:"shield_uid"`
	BusinessID string `json:"business_id"`
	Scatter    struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"scatter"`
	MarqueeID  string `json:"marquee_id"`
	NoticeType int    `json:"notice_type"`
}

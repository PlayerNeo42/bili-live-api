package dto

type User struct {
	Face       string `json:"face"`
	FaceFrame  string `json:"face_frame"`
	GuardLevel int    `json:"guard_level"`
	IsMainVip  int    `json:"is_main_vip"`
	IsSVIP     int    `json:"is_svip"`
	IsVip      int    `json:"is_vip"`
	LevelColor string `json:"level_color"`
	Manager    int    `json:"manager"`
	NameColor  string `json:"name_color"`
	Title      string `json:"title"`
	Uname      string `json:"uname"`
	UserLevel  int    `json:"user_level"`
}

type SuperChat struct {
	BackgroundBottomColor string  `json:"background_bottom_color"`
	BackgroundColor       string  `json:"background_color"`
	BackgroundColorEnd    string  `json:"background_color_end"`
	BackgroundColorStart  string  `json:"background_color_start"`
	BackgroundIcon        string  `json:"background_icon"`
	BackgroundImage       string  `json:"background_image"`
	BackgroundPriceColor  string  `json:"background_price_color"`
	ColorPoint            float64 `json:"color_point"`
	DmScore               int     `json:"dmscore"`
	EndTime               int     `json:"end_time"`
	Gift                  struct {
		GiftID   int    `json:"gift_id"`
		GiftName string `json:"gift_name"`
		Num      int    `json:"num"`
	} `json:"gift"`
	ID               int    `json:"id"`
	IsRanked         int    `json:"is_ranked"`
	IsSendAudit      int    `json:"is_send_audit"`
	MedalInfo        Medal  `json:"medal_info"`
	Message          string `json:"message"`
	MessageFontColor string `json:"message_font_color"`
	MessageTrans     string `json:"message_trans"`
	Price            int    `json:"price"`
	Rate             int    `json:"rate"`
	StartTime        int    `json:"start_time"`
	Time             int    `json:"time"`
	Token            string `json:"token"`
	TransMark        int    `json:"trans_mark"`
	Ts               int    `json:"ts"`
	Uid              int    `json:"uid"`
	UserInfo         User   `json:"user_info"`
}

type SuperChatDelete struct {
	IDs []int `json:"ids"`
}

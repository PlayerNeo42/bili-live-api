package dto

// SuperChat 超级留言(SC)
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
	ID               int      `json:"id"`
	IsRanked         int      `json:"is_ranked"`
	IsSendAudit      int      `json:"is_send_audit"`
	MedalInfo        Medal    `json:"medal_info"`
	Message          string   `json:"message"`
	MessageFontColor string   `json:"message_font_color"`
	MessageTrans     string   `json:"message_trans"`
	Price            int      `json:"price"`
	Rate             int      `json:"rate"`
	StartTime        int      `json:"start_time"`
	Time             int      `json:"time"`
	Token            string   `json:"token"`
	TransMark        int      `json:"trans_mark"`
	Ts               int      `json:"ts"`
	UID              int      `json:"uid"`
	UserInfo         UserInfo `json:"user_info"`
}

// SuperChatDelete 删除SC,IDs数组中的ID对应superChat.ID
type SuperChatDelete struct {
	IDs []int `json:"ids"`
}

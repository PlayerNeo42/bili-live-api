package dto

// AnchorLotAward 天选之人中奖完整信息
type AnchorLotAward struct {
	AwardDontPopup int    `json:"award_dont_popup"`
	AwardImage     string `json:"award_image"`
	AwardName      string `json:"award_name"`
	AwardNum       int    `json:"award_num"`
	AwardUsers     []struct {
		UID   int    `json:"uid"`
		Uname string `json:"uname"`
		Face  string `json:"face"`
		Level int    `json:"level"`
		Color int    `json:"color"`
	} `json:"award_users"`
	ID        int    `json:"id"`
	LotStatus int    `json:"lot_status"`
	URL       string `json:"url"`
	WebURL    string `json:"web_url"`
}

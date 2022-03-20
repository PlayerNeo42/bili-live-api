package dto

// AnchorLotStart 天选之人抽奖开始
type AnchorLotStart struct {
	AssetIcon      string `json:"asset_icon"`
	AwardImage     string `json:"award_image"`
	AwardName      string `json:"award_name"`
	AwardNum       int    `json:"award_num"`
	CurGiftNum     int    `json:"cur_gift_num"`
	CurrentTime    int    `json:"current_time"`
	Danmu          string `json:"danmu"`
	GiftID         int    `json:"gift_id"`
	GiftName       string `json:"gift_name"`
	GiftNum        int    `json:"gift_num"`
	GiftPrice      int    `json:"gift_price"`
	GoawayTime     int    `json:"goaway_time"`
	GoodsID        int    `json:"goods_id"`
	ID             int    `json:"id"`
	IsBroadcast    int    `json:"is_broadcast"`
	JoinType       int    `json:"join_type"`
	LotStatus      int    `json:"lot_status"`
	MaxTime        int    `json:"max_time"`
	RequireText    string `json:"require_text"`
	RequireType    int    `json:"require_type"`
	RequireValue   int    `json:"require_value"`
	RoomId         int    `json:"room_id"`
	SendGiftEnsure int    `json:"send_gift_ensure"`
	ShowPanel      int    `json:"show_panel"`
	StartDontPopup int    `json:"start_dont_popup"`
	Status         int    `json:"status"`
	Time           int    `json:"time"`
	URL            string `json:"url"`
	WebURL         string `json:"web_url"`
}

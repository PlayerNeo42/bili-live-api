package dto

type Gift struct {
	Action            string      `json:"action"`
	BatchComboID      string      `json:"batch_combo_id"`
	BatchComboSend    interface{} `json:"batch_combo_send"`
	BeatID            string      `json:"beatId"`
	BizSource         string      `json:"biz_source"`
	BlindGift         interface{} `json:"blind_gift"`
	BroadcastID       int         `json:"broadcast_id"`
	CoinType          string      `json:"coin_type"` // gold或者silver
	ComboResourcesID  int         `json:"combo_resources_id"`
	ComboSend         interface{} `json:"combo_send"`
	ComboStayTime     int         `json:"combo_stay_time"`
	ComboTotalCoin    int         `json:"combo_total_coin"`
	CritProb          int         `json:"crit_prob"`
	Demarcation       int         `json:"demarcation"`
	DiscountPrice     int         `json:"discount_price"`
	DmScore           int         `json:"dmscore"`
	Draw              int         `json:"draw"`
	Effect            int         `json:"effect"`
	EffectBlock       int         `json:"effect_block"`
	Face              string      `json:"face"`
	FloatScResourceID int         `json:"float_sc_resource_id"`
	GiftID            int         `json:"giftId"`
	GiftName          string      `json:"giftName"`
	GiftType          int         `json:"giftType"`
	Gold              int         `json:"gold"`
	GuardLevel        int         `json:"guard_level"`
	IsFirst           bool        `json:"is_first"`
	IsSpecialBatch    int         `json:"is_special_batch"`
	Magnification     float64     `json:"magnification"`
	MedalInfo         Medal       `json:"medal_info"`
	NameColor         string      `json:"name_color"`
	Num               int         `json:"num"`
	OriginalGiftName  string      `json:"original_gift_name"`
	Price             int         `json:"price"` // 除以1000后为实际价格(元)
	Rcost             int         `json:"rcost"`
	Remain            int         `json:"remain"`
	Rnd               string      `json:"rnd"`
	SendMaster        interface{} `json:"send_master"`
	Silver            int         `json:"silver"`
	Super             int         `json:"super"`
	SuperBatchGiftNum int         `json:"super_batch_gift_num"`
	SuperGiftNum      int         `json:"super_gift_num"`
	SvgaBlock         int         `json:"svga_block"`
	TagImage          string      `json:"tag_image"`
	Tid               string      `json:"tid"`
	Timestamp         int         `json:"timestamp"`
	TopList           interface{} `json:"top_list"`
	TotalCoin         int         `json:"total_coin"`
	UID               int         `json:"uid"`
	Uname             string      `json:"uname"`
}

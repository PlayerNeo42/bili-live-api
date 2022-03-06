package dto

type ComboSend struct {
	Action         string      `json:"action"`
	BatchComboID   string      `json:"batch_combo_id"`
	BatchComboNum  int         `json:"batch_combo_num"`
	ComboID        string      `json:"combo_id"`
	ComboNum       int         `json:"combo_num"`
	ComboTotalCoin int         `json:"combo_total_coin"`
	DmScore        int         `json:"dmscore"`
	GiftID         int         `json:"gift_id"`
	GiftName       string      `json:"gift_name"`
	GiftNum        int         `json:"gift_num"`
	IsShow         int         `json:"is_show"`
	MedalInfo      Medal       `json:"medal_info"`
	NameColor      string      `json:"name_color"`
	RUname         string      `json:"r_uname"`
	Ruid           int         `json:"ruid"`
	SendMaster     interface{} `json:"send_master"`
	TotalNum       int         `json:"total_num"`
	UID            int         `json:"uid"`
	Uname          string      `json:"uname"`
}

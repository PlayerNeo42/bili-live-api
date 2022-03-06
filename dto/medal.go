package dto

type Medal struct {
	AnchorRoomID     int    `json:"anchor_roomid"`
	AnchorUname      string `json:"anchor_uname"`
	GuardLevel       int    `json:"guard_level"`
	IconID           int    `json:"icon_id"`
	IsLighted        int    `json:"is_lighted"`
	MedalColor       int    `json:"medal_color"`
	MedalColorBorder int    `json:"medal_color_border"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalLevel       int    `json:"medal_level"`
	MedalName        string `json:"medal_name"`
	Special          string `json:"special"`
	TargetID         int    `json:"target_id"`
}

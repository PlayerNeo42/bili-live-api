package dto

type PopularityRedPocketStart struct {
	LotId           int    `json:"lot_id"`
	SenderUid       int    `json:"sender_uid"`
	SenderName      string `json:"sender_name"`
	SenderFace      string `json:"sender_face"`
	JoinRequirement int    `json:"join_requirement"`
	Danmu           string `json:"danmu"`
	CurrentTime     int    `json:"current_time"`
	StartTime       int    `json:"start_time"`
	EndTime         int    `json:"end_time"`
	LastTime        int    `json:"last_time"`
	RemoveTime      int    `json:"remove_time"`
	ReplaceTime     int    `json:"replace_time"`
	LotStatus       int    `json:"lot_status"`
	H5Url           string `json:"h5_url"`
	UserStatus      int    `json:"user_status"`
	Awards          []struct {
		GiftID   int    `json:"gift_id"`
		GiftName string `json:"gift_name"`
		GiftPic  string `json:"gift_pic"`
		Num      int    `json:"num"`
	} `json:"awards"`
	LotConfigId int `json:"lot_config_id"`
	TotalPrice  int `json:"total_price"`
}

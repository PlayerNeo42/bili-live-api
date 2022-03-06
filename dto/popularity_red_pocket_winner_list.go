package dto

type PopularityRedPocketWinnerList struct {
	LotId      int `json:"lot_id"`
	TotalNum   int `json:"total_num"`
	WinnerInfo []struct {
		Uid         int    `json:"uid"`
		Name        string `json:"name"`
		UserType    int    `json:"user_type"`
		AwardType   int    `json:"award_type"`
		AwardId     int    `json:"award_id"`
		AwardName   string `json:"award_name"`
		AwardPic    string `json:"award_pic"`
		AwardBigPic string `json:"award_big_pic"`
		AwardPrice  int    `json:"award_price"`
		BagId       int    `json:"bag_id"`
		GiftId      int    `json:"gift_id"`
		GiftNum     int    `json:"gift_num"`
	} `json:"winner_info"`
}

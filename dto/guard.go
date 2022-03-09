package dto

type Guard struct {
	UID        int    `json:"uid"`
	Username   string `json:"username"`
	GuardLevel int    `json:"guard_level"` // 3: 舰长, 2: 提督, 1: 总督
	Num        int    `json:"num"`
	Price      int    `json:"price"`
	GiftID     int    `json:"gift_id"`
	GiftName   string `json:"gift_name"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
}

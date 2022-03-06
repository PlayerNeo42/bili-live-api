package dto

type FansChange struct {
	RoomID    int `json:"roomid"`
	Fans      int `json:"fans"`
	RedNotice int `json:"red_notice"`
	FansClub  int `json:"fans_club"`
}

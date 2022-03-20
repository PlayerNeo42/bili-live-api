package dto

// RoomBlockMsg 房管禁言
type RoomBlockMsg struct {
	DmScore  int    `json:"dmscore"`
	Operator int    `json:"operator"`
	UID      int    `json:"uid"`
	Uname    string `json:"uname"`
}

package dto

// RoomChange 疑似房间标题变更
type RoomChange struct {
	Title          string `json:"title"`
	AreaID         int    `json:"area_id"`
	ParentAreaID   int    `json:"parent_area_id"`
	AreaName       string `json:"area_name"`
	ParentAreaName string `json:"parent_area_name"`
	LiveKey        string `json:"live_key"`
	SubSessionKey  string `json:"sub_session_key"`
}

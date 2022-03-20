package dto

// MatchRoomConf 未知
type MatchRoomConf struct {
	Type        string `json:"type"`
	CloseButton string `json:"close_button"`
	ForcePush   string `json:"force_push"`
	ButtonName  string `json:"button_name"`
	Background  string `json:"background"`
	ConfID      string `json:"conf_id"`
	ConfName    string `json:"conf_name"`
	RoomsInfo   []struct {
		RoomID     string `json:"room_id"`
		RoomName   string `json:"room_name"`
		LiveStatus int    `json:"live_status"`
	} `json:"rooms_info"`
	SeasonInfo    []interface{} `json:"season_info"`
	BackgroundURL string        `json:"background_url"`
	Scatter       struct {
		Max int `json:"max"`
		Min int `json:"min"`
	} `json:"scatter"`
	ButtonLink string `json:"button_link"`
	RoomsColor struct {
		FontColor       string `json:"font_color"`
		BackgroundColor string `json:"background_color"`
		BorderColor     string `json:"border_color"`
	} `json:"rooms_color"`
	State int `json:"state"`
}

package dto

type UserInfo struct {
	Face       string `json:"face"`
	FaceFrame  string `json:"face_frame"`
	GuardLevel int    `json:"guard_level"`
	IsMainVip  int    `json:"is_main_vip"`
	IsSVIP     int    `json:"is_svip"`
	IsVip      int    `json:"is_vip"`
	LevelColor string `json:"level_color"`
	Manager    int    `json:"manager"`
	NameColor  string `json:"name_color"`
	Title      string `json:"title"`
	Uname      string `json:"uname"`
	UserLevel  int    `json:"user_level"`
}

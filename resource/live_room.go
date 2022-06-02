package resource

import (
	"strconv"
)

// RoomInfoResp 直播房间信息
type RoomInfoResp struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		RoomID      int   `json:"room_id"`
		ShortID     int   `json:"short_id"`
		UID         int   `json:"uid"`
		NeedP2P     int   `json:"need_p2p"`
		IsHidden    bool  `json:"is_hidden"`
		IsLocked    bool  `json:"is_locked"`
		IsPortrait  bool  `json:"is_portrait"`
		LiveStatus  int   `json:"live_status"`
		HiddenTill  int   `json:"hidden_till"`
		LockTill    int   `json:"lock_till"`
		Encrypted   bool  `json:"encrypted"`
		PwdVerified bool  `json:"pwd_verified"`
		LiveTime    int64 `json:"live_time"`
		RoomShield  int   `json:"room_shield"`
		IsSp        int   `json:"is_sp"`
		SpecialType int   `json:"special_type"`
	} `json:"data"`
}

// GetRoomInfo 获取直播间详细信息
func GetRoomInfo(shortID int) (*RoomInfoResp, error) {
	result := &RoomInfoResp{}
	_, err := liveAPIClient.R().
		EnableTrace().
		SetQueryParam("id", strconv.Itoa(shortID)).
		SetResult(result).
		Get("/room/v1/Room/room_init")
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RealRoomID 获取真实直播房间号,主要用于websocket连接
func RealRoomID(shortID int) (int, error) {
	info, err := GetRoomInfo(shortID)
	if err != nil {
		return 0, err
	}
	return info.Data.RoomID, nil
}

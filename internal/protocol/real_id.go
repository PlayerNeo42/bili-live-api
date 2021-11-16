package protocol

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type RoomInfo struct {
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

func GetRealRoomId(roomId int) int {
	url := `https://api.live.bilibili.com/room/v1/Room/room_init?id=` + strconv.Itoa(roomId)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	room := RoomInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&room); err != nil {
		log.Fatal(err)
	}
	return room.Data.RoomID
}

package protocol

import (
	"encoding/json"
)

type EnterRoom struct {
	UID       int    `json:"uid"`
	RoomID    int    `json:"roomid"`
	ProtoVer  int    `json:"protover"`
	Platform  string `json:"platform"`
	ClientVer string `json:"clientver"`
}

func NewEnterRoom(realRoomID int) *EnterRoom {
	return &EnterRoom{
		UID:       0,
		RoomID:    realRoomID,
		ProtoVer:  1,
		Platform:  "web",
		ClientVer: "1.6.3",
	}
}

func (e *EnterRoom) JSON() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		log.Fatalln("impossible error", err)
	}
	return marshal
}

package dto

type WSPayload struct {
	RawBytes        []byte // RawBytes is the encoded result
	PacketLength    int    // PacketLength will be calculated during Encode()
	HeaderLength    int    // HeaderLength is currently a fixed number 16
	ProtocolVersion ProtoVer
	Operation       OPCode
	SequenceID      int
	Body            []byte
}

type WSNotificationBody struct {
	Cmd    string      `json:"cmd"`
	Data   interface{} `json:"data"`
	RoomID int         `json:"roomid"`
}

type WSEnterRoomBody struct {
	UID       int    `json:"uid"`
	RoomID    int    `json:"roomid"`
	ProtoVer  int    `json:"protover"`
	Platform  string `json:"platform"`
	ClientVer string `json:"clientver"`
}

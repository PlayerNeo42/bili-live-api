package dto

// OPCode The WebSocket OpCode.
type OPCode int

// WS opcodes
const (
	_ OPCode = iota
	_
	Heartbeat
	HeartbeatResponse
	_
	Notification
	_
	RoomEnter
	RoomEnterResponse
)

package api

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/resource"
	"github.com/botplayerneo/bili-live-api/websocket"
)

// Live 使用 NewLive() 来初始化
type Live struct {
	client *websocket.Client
	roomID int
}

// NewLive 构造函数
func NewLive(roomID int) *Live {
	return &Live{
		roomID: roomID,
		client: websocket.New(),
	}
}

// Start 接收房间号，开始websocket心跳连接并阻塞
func (l *Live) Start() {
	id, err := resource.GetRoomID(l.roomID)
	if err != nil {
		log.Errorf("获取房间ID失败：%v", err)
		return
	}

	if err := l.client.Connect(); err != nil {
		log.Errorf("连接websocket失败：%v", err)
		return
	}

	go l.enterRoom(id)

	if err := l.client.Listening(); err != nil {
		log.Errorf("监听websocket失败：%v", err)
		return
	}
}

func (l *Live) RegisterHandlers(handlers ...interface{}) {
	websocket.RegisterHandlers(handlers...)
}

// 发送进入房间请求
func (l *Live) enterRoom(id int) {
	log.Infof("进入房间：%d", id)
	// 忽略错误
	var err error
	body, _ := jsoniter.Marshal(dto.WSEnterRoomBody{
		RoomID:    id, // 真实房间ID
		ProtoVer:  1,  // 填1
		Platform:  "web",
		ClientVer: "1.6.3",
	})
	if err = l.client.Write(&dto.WSPayload{
		ProtocolVersion: dto.JSON,
		Operation:       dto.RoomEnter,
		Body:            body,
	}); err != nil {
		log.Errorf("发送进入房间请求失败：%v", err)
		return
	}
}
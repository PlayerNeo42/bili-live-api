package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/spellingDragon/bili-live-api/dto"
	"github.com/spellingDragon/bili-live-api/log"
	"github.com/spellingDragon/bili-live-api/resource"
	"github.com/spellingDragon/bili-live-api/websocket"
)

// Live 使用 NewLive() 来初始化
type Live struct {
	Client     *websocket.Client
	RoomID     int
	LiverUname string
	LastTitle  string
	Face       string
}

// NewLive 构造函数
func NewLive(roomID int) *Live {
	return &Live{
		Client: websocket.New(),
		RoomID: roomID,
	}
}

// Start 接收房间号，开始websocket心跳连接并阻塞
func (l *Live) Start() {
	for {
		if err := l.Listen(); err != nil {
			if strings.Contains(err.Error(), "获取房间号失败") {
				l.Stop()
				break
			}
			log.Warnf("连接失败, 重连中...:%v", err)
			time.Sleep(10 * time.Second)
			if err := l.Client.Connect(); err != nil {
				log.Errorf("连接websocket失败：%v", err)
				return
			}
			continue
		}
		break
	}
}

// Start 接收房间号，开始websocket心跳连接并阻塞
func (l *Live) Stop() {
	l.Client.Close()
}

func (l *Live) Listen() error {
	roomInfoResponse, err := resource.GetRoomInfo(l.RoomID)
	if err != nil {
		return fmt.Errorf("获取房间号失败：%v", err)
	}

	if err := l.Client.Connect(); err != nil {
		return fmt.Errorf("连接websocket失败：%v", err)
	}

	// TODO 发送进房包,可能有顺序问题
	go l.enterRoom(roomInfoResponse)

	if err := l.Client.Listening(); err != nil {
		return fmt.Errorf("监听websocket失败：%v", err)
	}
	return nil
}

// RegisterHandlers 注册不同的事件处理
// handler类型需要是定义在 websocket/handler_registration.go 中的类型，如:
// - websocket.DanmakuHandler
// - websocket.GiftHandler
// - websocket.GuardHandler
func (l *Live) RegisterHandlers(handlers ...interface{}) error {
	return l.Client.RegisterHandlers(handlers...)
}

// 发送进入房间请求
func (l *Live) enterRoom(roomInfo *resource.RoomInfoResp) {
	roomInfoJson, _ := json.Marshal(roomInfo)
	log.Infof("进入房间：%s", string(roomInfoJson))
	liverInfo, _ := resource.UserInfo(roomInfo.Data.UID)
	liverInfoJson, _ := json.Marshal(liverInfo)
	log.Infof("主播信息：%s", string(liverInfoJson))
	l.setLiverProfile(liverInfo)
	// 忽略错误
	var err error
	body, _ := jsoniter.Marshal(dto.WSEnterRoomBody{
		RoomID:   roomInfo.Data.RoomID, // 真实房间ID
		ProtoVer: 3,                    // 填3
		Platform: "web",
		Type:     2,
	})
	if err = l.Client.Write(&dto.WSPayload{
		ProtocolVersion: dto.JSON,
		Operation:       dto.RoomEnter,
		Body:            body,
	}); err != nil {
		log.Errorf("发送进入房间请求失败：%v", err)
		return
	}
}

func (l *Live) RefreshRoom() error {
	roomInfo, err := resource.GetRoomInfo(l.RoomID)
	if err != nil {
		return fmt.Errorf("刷新房间信息失败：%v", err)
	}
	liverInfo, _ := resource.UserInfo(roomInfo.Data.UID)
	if err != nil {
		return fmt.Errorf("刷新主播信息失败：%v", err)
	}
	liverInfoJson, _ := json.Marshal(liverInfo)
	log.Infof("主播信息：%s", string(liverInfoJson))
	l.setLiverProfile(liverInfo)
	return nil
}

func (l *Live) setLiverProfile(liverInfo *resource.UserInfoResp) {
	l.LiverUname = liverInfo.Data.Name
	l.LastTitle = liverInfo.Data.LiveRoom.Title
	l.Face = liverInfo.Data.Face
}

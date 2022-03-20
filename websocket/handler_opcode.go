package websocket

import (
	"encoding/binary"

	jsoniter "github.com/json-iterator/go"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
)

type opCodeHandler func(*dto.WSPayload)

var opCodeHandlerMap = map[dto.OPCode]opCodeHandler{
	dto.Notification:      notificationHandler,
	dto.HeartbeatResponse: heartbeatResponseHandler,
	dto.RoomEnterResponse: roomEnterResponseHandler,
}

// 通知类消息，弹幕、礼物等
func notificationHandler(payload *dto.WSPayload) {
	eventType := dto.EventType(jsoniter.Get(payload.Body, "cmd").ToString())
	handler, ok := eventPayloadHandlerMap[eventType]
	if !ok {
		log.Debugf("未知cmd类型: %s", eventType)
		unknownEventHandler(payload)
		return
	}
	handler(payload)
}

// 心跳回应,body为人气值
func heartbeatResponseHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Popularity == nil {
		return
	}
	popularity := binary.BigEndian.Uint32(payload.Body)
	DefaultEventHandlers.Popularity(popularity)
}

// 进房回应，body为空
func roomEnterResponseHandler(payload *dto.WSPayload) {
	log.Info("进房成功")
}

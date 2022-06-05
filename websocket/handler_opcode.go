package websocket

import (
	"encoding/binary"
	"strings"

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
	eType := jsoniter.Get(payload.Body, "cmd").ToString()
	var handler eventPayloadHandler
	// HACK 更新后收到的cmd会变为"DANMU_MSG:4:0:2:2:2:0", 在此特殊处理
	if strings.HasPrefix(eType, "DANMU_MSG") {
		handler = eventPayloadHandlerMap[dto.EventDanmaku]
	} else {
		eventType := dto.EventType(eType)
		log.Debug("收到CMD消息:", eventType, string(payload.Body))
		var ok bool
		handler, ok = eventPayloadHandlerMap[eventType]
		if !ok {
			log.Debugf("未知cmd类型: %s", eventType)
			handler = unknownEventHandler
		}
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
	log.Debug("收到心跳回应,人气值:", popularity)
}

// 进房回应，body为空
func roomEnterResponseHandler(payload *dto.WSPayload) {
	log.Debug("进房成功")
}

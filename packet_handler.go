package api

import (
	"github.com/agentanderson/bili-live-api/internal/protocol"
	"github.com/agentanderson/bili-live-api/message"
	jsoniter "github.com/json-iterator/go"
)

func danmakuHandler(packet protocol.Packet) message.Danmaku {
	raw := jsoniter.Get(packet.Body, "info")
	return message.Danmaku{
		Uid:       raw.Get(2, 0).ToInt(),
		Uname:     raw.Get(2, 1).ToString(),
		Content:   raw.Get(1).ToString(),
		Timestamp: raw.Get(0, 4).ToInt(),
	}
}

func superChatHandler(packet protocol.Packet) message.SuperChat {
	raw := jsoniter.Get(packet.Body, "data")
	return message.SuperChat{
		Uid:     raw.Get("uid").ToInt(),
		Uname:   raw.Get("user_info", "uname").ToString(),
		Message: raw.Get("message").ToString(),
		Price:   raw.Get("price").ToInt(),
	}
}

func giftHandler(packet protocol.Packet) message.Gift {
	raw := jsoniter.Get(packet.Body, "data")
	return message.Gift{
		Uid:      raw.Get("uid").ToInt(),
		Uname:    raw.Get("uname").ToString(),
		Price:    raw.Get("price").ToInt(),
		GiftName: raw.Get("giftName").ToString(),
	}
}

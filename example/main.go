package main

import (
	"fmt"

	api "github.com/botplayerneo/bili-live-api"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/websocket"
)

func main() {
	l := api.NewLive(213)

	l.RegisterHandlers(
		danmakuHandler(),
		giftHandler(),
		guardHandler(),
		superChatHandler(),
	)

	l.Start()
}

func danmakuHandler() websocket.DanmakuHandler {
	return func(danmaku *dto.Danmaku) {
		fmt.Printf("%s:%s\n", danmaku.Uname, danmaku.Content)
	}
}

func giftHandler() websocket.GiftHandler {
	return func(gift *dto.Gift) {
		fmt.Printf("%s 赠送 [礼物%.1f￥]%s x %d\n", gift.Uname, float64(gift.Price)/1000, gift.GiftName, gift.Num)
	}
}

func guardHandler() websocket.GuardHandler {
	return func(guard *dto.Guard) {
		fmt.Printf("%s 开通舰长,级别%d", guard.Username, guard.GuardLevel)
	}
}

func superChatHandler() websocket.SuperChatHandler {
	return func(superChat *dto.SuperChat) {
		fmt.Printf("[%d￥SC]%s:%s\n", superChat.Price, superChat.UserInfo.Uname, superChat.Message)
	}
}

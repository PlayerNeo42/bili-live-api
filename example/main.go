package main

import (
	"fmt"

	api "github.com/botplayerneo/bili-live-api"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/websocket"
)

func main() {
	// 设置内部 logger 打印等级
	log.SetLogLevel(log.LevelWarn)

	// 参数为房间号
	l := api.NewLive(510)

	// 内部通过type switch来判断handler类型
	// handler类型定义在websocket包中
	err := l.RegisterHandlers(
		danmakuHandler(),
		giftHandler(),
		guardHandler(),
		superChatHandler(),
	)
	if err != nil {
		panic(err)
	}

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

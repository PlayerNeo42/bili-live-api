# B站直播API

B站直播相关API的Go语言实现

目前仅支持直播间弹幕相关API

## 支持

- 礼物
- 超级留言
- 弹幕

详见`dto/`目录下的文件

## 快速开始

### 安装

```shell 
go get github.com/botplayerneo/bili-live-api
```

### 使用

```go
package main

import (
	"fmt"
	api "github.com/botplayerneo/bili-live-api"
	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/websocket"
)

func main() {
	// 参数为房间号
	l := api.NewLive(510)

	// 内部通过type switch来判断handler类型
	// handler类型定义在websocket包中
	l.RegisterHandlers(
		danmakuHandler(),
		giftHandler(),
		guardHandler(),
		superChatHandler(),
	)

	// 启动并阻塞协程
	l.Start()
}

func danmakuHandler() websocket.DanmakuHandler {
	return func(danmaku *dto.Danmaku) {
		fmt.Printf("%s:%s\n",
			danmaku.Uname,
			danmaku.Content,
		)
	}
}

func giftHandler() websocket.GiftHandler {
	return func(gift *dto.Gift) {
		fmt.Printf("%s 赠送 [礼物%.1f￥]%s x %d\n",
			gift.Uname,
			float64(gift.Price)/1000,
			gift.GiftName,
			gift.Num,
		)
	}
}

func guardHandler() websocket.GuardHandler {
	return func(guard *dto.Guard) {
		fmt.Printf("%s 开通舰长,级别%d",
			guard.Username,
			guard.GuardLevel,
		)
	}
}

func superChatHandler() websocket.SuperChatHandler {
	return func(superChat *dto.SuperChat) {
		fmt.Printf("[%d￥SC]%s:%s\n",
			superChat.Price,
			superChat.UserInfo.Uname,
			superChat.Message,
		)
	}
}
```

## 相关

项目组织结构参考: [botgo](https://github.com/tencent-connect/botgo)

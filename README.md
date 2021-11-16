# B站直播API

B站直播相关API的Go语言实现

目前仅支持直播间弹幕相关API

## 支持

- 弹幕接收和发送
- 礼物
- 超级留言

## Quick Start

### 安装

```shell 
go get github.com/agentanderson/bili-live-api
```

### 使用

```go
package main

import (
	"fmt"
	api "github.com/agentanderson/bili-live-api"
	"github.com/agentanderson/bili-live-api/message"
)

func main() {
	l := api.NewLive()

	l.OnDanmaku(func(danmaku message.Danmaku) {
		fmt.Printf("%s:%s", danmaku.Uname, danmaku.Content)
	})

	l.OnGift(func(gift message.Gift) {
		if gift.Price == 0 {
			return
		}
		fmt.Printf("%+v\n", gift)
	})

	l.Start(510)
}
``
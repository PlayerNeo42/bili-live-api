package api

import (
	"fmt"
	"github.com/agentanderson/bili-live-api/internal/protocol"
	"github.com/agentanderson/bili-live-api/message"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var log = protocol.GetLogger()

// Live 使用 NewLive() 来初始化
type Live struct {
	client      *protocol.Client
	roomID      int
	realRoomID  int
	credential  *Credential
	onDanmaku   func(danmaku message.Danmaku)
	onGift      func(gift message.Gift)
	onSuperChat func(superChat message.SuperChat)
}

// NewLive 构造函数，初始化回调函数
func NewLive() *Live {
	return &Live{
		onDanmaku:   func(danmaku message.Danmaku) {},
		onGift:      func(gift message.Gift) {},
		onSuperChat: func(superChat message.SuperChat) {},
	}
}

// Start 接收房间号，开始websocket心跳连接并阻塞当前goroutine
func (b *Live) Start(roomID int) {
	b.roomID = roomID
	realRoomID := protocol.GetRealRoomId(roomID)
	b.realRoomID = realRoomID
	b.client = protocol.NewClient()
	b.client.OnMessage(func(packet protocol.Packet) {
		switch packet.Operation {
		case protocol.Notification:
			cmd := jsoniter.Get(packet.Body, "cmd").ToString()
			switch cmd {
			case "DANMU_MSG":
				b.onDanmaku(danmakuHandler(packet))
			case "SUPER_CHAT_MESSAGE":
				b.onSuperChat(superChatHandler(packet))
			case "SEND_GIFT":
				b.onGift(giftHandler(packet))
			 //case "LIVE_INTERACTIVE_GAME":
			 //case "INTERACT_WORD":
			 //case "COMBO_SEND":
			 //case "ENTRY_EFFECT":
			 //case "WIDGET_BANNER":
			 //case "ONLINE_RANK_COUNT":
			 //case "ONLINE_RANK_V2":
			 //case "STOP_LIVE_ROOM_LIST":
			 //case "USER_TOAST_MSG":
			 //case "GUARD_BUY":
			 //case "NOTICE_MSG":
			 //case "ROOM_REAL_TIME_MESSAGE_UPDATE":
			 //case "ONLINE_RANK_TOP3":
			 //case "HOT_RANK_CHANGED":
			 //default:
			 //	log.WithField("data", string(packet.Body)).Warn("未知cmd")
			}
		//case protocol.HeartBeatResponse:
			//default:
			//	log.WithField("protover", packet.ProtocolVersion).
			//		WithField("data", string(packet.Body)).
			//		Warn("未知protover")
		}
	})
	enterRoomPacket := protocol.NewJSONPacket(protocol.RoomEnter, protocol.NewEnterRoom(realRoomID).JSON())
	if err := b.client.Send(enterRoomPacket); err != nil {
		log.Fatal(err)
	}
	b.client.StartHeartBeat()
}

// OnDanmaku 注册弹幕消息回调
func (b *Live) OnDanmaku(handler func(danmaku message.Danmaku)) {
	b.onDanmaku = handler
}

// OnGift 注册礼物消息回调
func (b *Live) OnGift(handler func(gift message.Gift)) {
	b.onGift = handler
}

// OnSuperChat 注册超级留言消息回调
func (b *Live) OnSuperChat(handler func(superChat message.SuperChat)) {
	b.onSuperChat = handler
}

// SetCredential 设置身份信息
func (b *Live) SetCredential(c *Credential) {
	b.credential = c
}

// SendDanmaku 发送弹幕，需要首先调用 Live.SetCredential 方法
func (b *Live) SendDanmaku(msg string) error {
	if b.credential == nil {
		panic("credential not found, use Live.SetCredential first")
	}

	// 默认使用白色
	color := "16777215"
	// color = "5566168"

	// 协议固定格式
	body := strings.NewReader(url.Values{
		"bubble":     {"0"},
		"msg":        {msg},
		"color":      {color},
		"mode":       {"1"},
		"fontsize":   {"25"},
		"rnd":        {strconv.FormatInt(time.Now().Unix(), 10)},
		"roomid":     {strconv.Itoa(b.realRoomID)},
		"csrf":       {b.credential.BiliJct},
		"csrf_token": {b.credential.BiliJct},
	}.Encode())
	req, err := http.NewRequest("POST", "https://api.live.bilibili.com/msg/send", body)
	if err != nil {
		panic(fmt.Errorf("impossible err: %w", err))
	}
	req.Header.Set("Authority", "api.live.bilibili.com")
	req.Header.Set("Sec-Ch-Ua", "\"Google Chrome\";v=\"95\", \"Chromium\";v=\"95\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://live.bilibili.com")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://live.bilibili.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")

	req.Header.Set("Cookie", b.credential.Cookie())

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("弹幕发送失败: %w", err)
	}
	return nil
}

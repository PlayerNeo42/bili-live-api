package protocol

import (
	"github.com/agentanderson/bili-live-api/internal/util"
	"time"

	"github.com/gorilla/websocket"
)

// protocol version
const (
	_ = iota
	_
	HeartBeat
	HeartBeatResponse
	_
	Notification
	_
	RoomEnter
	RoomEnterResponse
)

// operation type
const (
	JSON = iota
	Popularity
	Zlib
	Brotli
)

var (
	HearBeatPacket = NewPacket(JSON, HeartBeat, nil)
)

type Client struct {
	conn *websocket.Conn
}

func NewClient() *Client {
	client := new(Client)
	err := client.connect()
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (c *Client) connect() error {
	conn, _, err := websocket.DefaultDialer.Dial("wss://broadcastlv.chat.bilibili.com/sub", nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Send(packet Packet) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, packet.Build())
}

func (c *Client) OnMessage(handler func(Packet)) {
	go func() {
		for {
			msgType, data, err := c.conn.ReadMessage()
			if err != nil {
				err := util.Retry(c.connect, 10)
				if err != nil {
					log.Fatal(err)
				}
				continue
			}
			if msgType != websocket.BinaryMessage {
				log.Error("收到非二进制消息类型", data)
				continue
			}
			for _, packet := range Decode(data).Parse() {
				handler(packet)
			}
		}
	}()
}

func (c *Client) StartHeartBeat() {
	for {
		if err := c.conn.WriteMessage(websocket.BinaryMessage, HearBeatPacket.Build()); err != nil {
			log.Fatal(err)
		}
		time.Sleep(30 * time.Second)
	}
}

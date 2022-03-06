package websocket

import (
	"time"

	wss "github.com/gorilla/websocket"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/resource"
)

var heartbeatPayload = &dto.WSPayload{
	ProtocolVersion: dto.JSON,
	Operation:       dto.Heartbeat,
	Body:            nil,
}

type Client struct {
	conn            *wss.Conn
	heartbeatTicker *time.Ticker
	closeCh         chan struct{}
	messageCh       chan *dto.WSPayload
}

func New() *Client {
	return &Client{
		heartbeatTicker: time.NewTicker(15 * time.Second),
		messageCh:       make(chan *dto.WSPayload, 10000),
		closeCh:         make(chan struct{}, 10),
	}
}

func (c *Client) Connect() error {
	var err error
	c.conn, _, err = wss.DefaultDialer.Dial(resource.WSUrl, nil)
	if err != nil {
		log.Errorf("websocket connect error: %v", err)
		return err
	}
	log.Info("websocket连接成功")
	return nil
}

// Listening 开始心跳和消息监听,阻塞协程
func (c *Client) Listening() error {
	go c.readMessage()
	go c.handleMessage()

	// 开始心跳
	for {
		select {
		case <-c.closeCh:
			log.Infof("收到关闭信号,关闭心跳")
			c.Close()
			return nil
		case <-c.heartbeatTicker.C:
			err := c.Write(heartbeatPayload)
			if err != nil {
				log.Errorf("发送心跳失败: %v", err)
				return err
			}
		}
	}
}

func (c *Client) Close() {
	if err := c.conn.Close(); err != nil {
		log.Errorf("%s, close conn err: %v", err)
	}
	c.heartbeatTicker.Stop()
}

func (c *Client) Write(payload *dto.WSPayload) error {
	err := c.conn.WriteMessage(wss.BinaryMessage, Encode(payload))
	if err != nil {
		log.Errorf("websocket send: %v", err)
		return err
	}
	return nil
}

func (c *Client) readMessage() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			close(c.messageCh)
			c.closeCh <- struct{}{}
			log.Errorf("websocket read error: %v", err)
			return
		}
		c.messageCh <- Decode(message)
	}
}

func (c *Client) handleMessage() {
	for data := range c.messageCh {
		if err := parseAndHandle(data); err != nil {
			log.Errorf("websocket handle message error: %v", err)
		}
	}
	log.Warnf("message chan closed")
}

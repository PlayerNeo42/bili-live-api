package websocket

import (
	"time"

	wss "github.com/gorilla/websocket"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/resource"
)

const (
	heartBeatInterval = time.Second * 30
)

var heartbeatPayload = &dto.WSPayload{
	ProtocolVersion: dto.JSON,
	Operation:       dto.Heartbeat,
	Body:            nil,
}

// Client websocket客户端实例
type Client struct {
	conn            *wss.Conn
	heartbeatTicker *time.Ticker
	closeCh         chan struct{}
	messageCh       chan *dto.WSPayload
}

// New 创建websocket客户端
func New() *Client {
	return &Client{
		heartbeatTicker: time.NewTicker(heartBeatInterval),
		messageCh:       make(chan *dto.WSPayload, 10), // TODO debug
		closeCh:         make(chan struct{}, 10),
	}
}

// Connect 连接到B站直播服务端
func (c *Client) Connect() error {
	var err error
	c.conn, _, err = wss.DefaultDialer.Dial(resource.WSUrl, nil)
	if err != nil {
		log.Errorf("websocket connect error: %v", err)
		return err
	}
	log.Debug("websocket连接成功")
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
			log.Infof("发送心跳成功")
		}
	}
}

// Close 关闭websocket连接,停止心跳
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
	log.Debugf("websocket send: %+v", payload)
	return nil
}

func (c *Client) readMessage() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Errorf("websocket read error: %v", err)
			close(c.messageCh)
			c.closeCh <- struct{}{}
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

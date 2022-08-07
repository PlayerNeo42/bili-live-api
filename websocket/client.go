package websocket

import (
	"fmt"
	"time"

	wss "github.com/gorilla/websocket"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
	"github.com/botplayerneo/bili-live-api/resource"
)

const (
	heartbeatInterval = time.Second * 30
)

var heartbeatPayload = &dto.WSPayload{
	ProtocolVersion: dto.JSON,
	Operation:       dto.Heartbeat,
	Body:            nil,
}

// Client websocket客户端实例
type Client struct {
	conn                 *wss.Conn
	heartbeatTicker      *time.Ticker
	closeCh              chan bool
	messageCh            chan *dto.WSPayload
	DefaultEventHandlers DefaultEventHandlers
}

// New 创建websocket客户端
func New() *Client {
	return &Client{
		heartbeatTicker: time.NewTicker(heartbeatInterval),
		messageCh:       make(chan *dto.WSPayload, 10),
		closeCh:         make(chan bool, 10),
	}
}

// Connect 连接到B站直播服务端
func (client *Client) Connect() error {
	var err error
	client.conn, _, err = wss.DefaultDialer.Dial(resource.WSUrl, nil)
	if err != nil {
		log.Errorf("websocket connect error: %v", err)
		return err
	}
	log.Debug("websocket连接成功")
	return nil
}

// Listening 开始心跳和消息监听,阻塞协程
func (client *Client) Listening() error {
	go client.readMessage()
	go client.handleMessage()

	// 开始心跳
	for {
		select {
		case isNormal := <-client.closeCh:
			client.stopConnection()
			if isNormal {
				return nil
			} else {
				return fmt.Errorf("websocket连接关闭")
			}
		case <-client.heartbeatTicker.C:
			err := client.Write(heartbeatPayload)
			if err != nil {
				log.Errorf("发送心跳失败: %v", err)
				return err
			}
			log.Debugf("发送心跳成功")
		}
	}
}

func (client *Client) Close() {
	client.closeCh <- true
}

// stopConnection 关闭websocket连接,停止心跳
func (client *Client) stopConnection() {
	if err := client.conn.Close(); err != nil {
		log.Errorf("%s, close conn err: %v", err)
	}
	client.conn = nil
}

func (client *Client) Write(payload *dto.WSPayload) error {
	if client.conn == nil {
		log.Warnf("bili client connection is null")
		return nil
	}
	err := client.conn.WriteMessage(wss.BinaryMessage, Encode(payload))
	if err != nil {
		log.Errorf("websocket send: %v", err)
		return err
	}
	log.Debugf("websocket send: %+v", payload)
	return nil
}

func (client *Client) readMessage() error {
	if client.conn == nil {
		log.Warnf("bili client connection is null")
		return nil
	}
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			log.Errorf("websocket read error: %v", err)
			client.closeCh <- false
			return err
		}
		client.messageCh <- Decode(message)
	}
}

func (client *Client) handleMessage() {
	for data := range client.messageCh {
		if err := parseAndHandle(data, client); err != nil {
			log.Errorf("websocket handle message error: %v", err)
		}
	}
	log.Warnf("message chan closed")
}

// RegisterHandlers 注册不同的事件处理
// handler类型需要是定义在 websocket/handler_registration.go 中的类型，如:
// - DanmakuHandler
// - GiftHandler
// - GuardHandler
func (client *Client) RegisterHandlers(handlers ...interface{}) error {

	for _, h := range handlers {
		switch handler := h.(type) {
		case DanmakuHandler:
			client.DefaultEventHandlers.Danmaku = handler
		case GiftHandler:
			client.DefaultEventHandlers.Gift = handler
		case SuperChatHandler:
			client.DefaultEventHandlers.SuperChat = handler
		case SuperChatDeleteHandler:
			client.DefaultEventHandlers.SuperChatDelete = handler
		case GuardHandler:
			client.DefaultEventHandlers.Guard = handler
		case PopularityHandler:
			client.DefaultEventHandlers.Popularity = handler
		case LiveHandler:
			client.DefaultEventHandlers.Live = handler
		case PreparingHandler:
			client.DefaultEventHandlers.Preparing = handler
		case EntryEffectHandler:
			client.DefaultEventHandlers.EntryEffect = handler
		case InteractWordHandler:
			client.DefaultEventHandlers.InteractWord = handler
		case ComboSendHandler:
			client.DefaultEventHandlers.ComboSend = handler
		case FansChangeHandler:
			client.DefaultEventHandlers.FansChange = handler
		case InteractiveGameHandler:
			client.DefaultEventHandlers.InteractiveGame = handler
		case OnlineRankCountHandler:
			client.DefaultEventHandlers.OnlineRankCount = handler
		case HotRankChangedHandler:
			client.DefaultEventHandlers.HotRankChanged = handler
		case HotRankChangedV2Handler:
			client.DefaultEventHandlers.HotRankChangedV2 = handler
		case HotRankSettlementHandler:
			client.DefaultEventHandlers.HotRankSettlement = handler
		case HotRankSettlementV2Handler:
			client.DefaultEventHandlers.HotRankSettlementV2 = handler
		case OnlineRankTop3Handler:
			client.DefaultEventHandlers.OnlineRankTop3 = handler
		case OnlineRankV2Handler:
			client.DefaultEventHandlers.OnlineRankV2 = handler
		case StopLiveRoomListHandler:
			client.DefaultEventHandlers.StopLiveRoomList = handler
		case WatchedChangeHandler:
			client.DefaultEventHandlers.WatchedChange = handler
		case WidgetBannerHandler:
			client.DefaultEventHandlers.WidgetBanner = handler
		case PopularityRedPocketStartHandler:
			client.DefaultEventHandlers.PopularityRedPocketStart = handler
		case PopularityRedPocketWinnerListHandler:
			client.DefaultEventHandlers.PopularityRedPocketWinnerList = handler
		case NoticeMsgHandler:
			client.DefaultEventHandlers.NoticeMsg = handler
		case UnknownEventHandler:
			client.DefaultEventHandlers.UnknownEvent = handler
		case AnchorLotAwardHandler:
			client.DefaultEventHandlers.AnchorLotAward = handler
		case UserToastMsgHandler:
			client.DefaultEventHandlers.UserToastMsg = handler
		case RoomChangeHandler:
			client.DefaultEventHandlers.RoomChange = handler
		case RoomBlockMsgHandler:
			client.DefaultEventHandlers.RoomBlockMsg = handler
		case MatchRoomConfHandler:
			client.DefaultEventHandlers.MatchRoomConf = handler
		case CommonNoticeDanmakuHandler:
			client.DefaultEventHandlers.CommonNoticeDanmaku = handler
		case AnchorLotCheckStatusHandler:
			client.DefaultEventHandlers.AnchorLotCheckstatus = handler
		case AnchorLotEndHandler:
			client.DefaultEventHandlers.AnchorLotEnd = handler
		case AnchorLotStartHandler:
			client.DefaultEventHandlers.AnchorLotStart = handler
		case TradingScoreHandler:
			client.DefaultEventHandlers.TradingScore = handler
		default:
			return fmt.Errorf("未知handler类型: %T", handler)
		}
	}
	return nil
}

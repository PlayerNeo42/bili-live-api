package websocket

import (
	"os"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
)

// LiveHandler 开播
type LiveHandler func()

// PreparingHandler 下播
type PreparingHandler func()

// DanmakuHandler 弹幕
type DanmakuHandler func(danmaku *dto.Danmaku)

// GiftHandler 礼物
type GiftHandler func(gift *dto.Gift)

// SuperChatHandler 超级留言
type SuperChatHandler func(superChat *dto.SuperChat)

// SuperChatDeleteHandler 删除超级留言
type SuperChatDeleteHandler func(superChat *dto.SuperChatDelete)

// GuardHandler 购买舰长
type GuardHandler func(guard *dto.Guard)

// PopularityHandler 人气变动
type PopularityHandler func(popularity uint32)

// EntryEffectHandler 进入房间特效
type EntryEffectHandler func(entryEffect *dto.EntryEffect)

// InteractWordHandler 发生互动
type InteractWordHandler func(interactWord *dto.InteractWord)

// ComboSendHandler 礼物连击
type ComboSendHandler func(comboSend *dto.ComboSend)

// FansChangeHandler 粉丝变动
type FansChangeHandler func(fansChange *dto.FansChange)

// InteractiveGameHandler 游戏互动
type InteractiveGameHandler func(interactiveGame *dto.InteractiveGame)

// OnlineRankCountHandler 在线榜单
type OnlineRankCountHandler func(onlineRankCount *dto.OnlineRankCount)

// HotRankChangedHandler 热度榜单变化
type HotRankChangedHandler func(hotRankChanged *dto.HotRankChanged)

// HotRankChangedV2Handler 热度榜单变化V2
type HotRankChangedV2Handler func(hotRankChangedV2 *dto.HotRankChangedV2)

// HotRankSettlementHandler 热度榜单结算
type HotRankSettlementHandler func(hotRankSettlement *dto.HotRankSettlement)

// HotRankSettlementV2Handler 热度榜单结算V2
type HotRankSettlementV2Handler func(hotRankSettlementV2 *dto.HotRankSettlementV2)

// OnlineRankTop3Handler 高能榜前三
type OnlineRankTop3Handler func(onlineRankTop3 *dto.OnlineRankTop3)

// OnlineRankV2Handler 高能榜V2
type OnlineRankV2Handler func(onlineRankV2 *dto.OnlineRankV2)

// StopLiveRoomListHandler 停止直播列表
type StopLiveRoomListHandler func(stopLiveRoomList *dto.StopLiveRoomList)

// WatchedChangeHandler 观看人数变化
type WatchedChangeHandler func(watchedChange *dto.WatchedChange)

// WidgetBannerHandler 推广横幅
type WidgetBannerHandler func(widgetBanner *dto.WidgetBanner)

// PopularityRedPocketStartHandler 红包开始
type PopularityRedPocketStartHandler func(popularityRedPocketStart *dto.PopularityRedPocketStart)

// PopularityRedPocketWinnerListHandler 红包中奖名单
type PopularityRedPocketWinnerListHandler func(popularityRedPocketWinnerList *dto.PopularityRedPocketWinnerList)

// NoticeMsgHandler 公告消息
type NoticeMsgHandler func(noticeMsg *dto.NoticeMsg)

// handler_type above(for hygen)

// DefaultEventHandlers 默认事件处理器,由 RegisterHandlers 注册
var DefaultEventHandlers struct {
	Live                          LiveHandler
	Preparing                     PreparingHandler
	Danmaku                       DanmakuHandler
	Gift                          GiftHandler
	SuperChat                     SuperChatHandler
	SuperChatDelete               SuperChatDeleteHandler
	Guard                         GuardHandler
	Popularity                    PopularityHandler
	EntryEffect                   EntryEffectHandler
	InteractWord                  InteractWordHandler
	ComboSend                     ComboSendHandler
	FansChange                    FansChangeHandler
	InteractiveGame               InteractiveGameHandler
	OnlineRankCount               OnlineRankCountHandler
	HotRankChanged                HotRankChangedHandler
	HotRankChangedV2              HotRankChangedV2Handler
	HotRankSettlement             HotRankSettlementHandler
	HotRankSettlementV2           HotRankSettlementV2Handler
	OnlineRankTop3                OnlineRankTop3Handler
	OnlineRankV2                  OnlineRankV2Handler
	StopLiveRoomList              StopLiveRoomListHandler
	WatchedChange                 WatchedChangeHandler
	WidgetBanner                  WidgetBannerHandler
	PopularityRedPocketStart      PopularityRedPocketStartHandler
	PopularityRedPocketWinnerList PopularityRedPocketWinnerListHandler
	NoticeMsg                     NoticeMsgHandler
	// handler_struct above(for hygen)
}

// RegisterHandlers 注册事件处理函数
func RegisterHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handler := h.(type) {
		case DanmakuHandler:
			DefaultEventHandlers.Danmaku = handler
		case GiftHandler:
			DefaultEventHandlers.Gift = handler
		case SuperChatHandler:
			DefaultEventHandlers.SuperChat = handler
		case SuperChatDeleteHandler:
			DefaultEventHandlers.SuperChatDelete = handler
		case GuardHandler:
			DefaultEventHandlers.Guard = handler
		case PopularityHandler:
			DefaultEventHandlers.Popularity = handler
		case EntryEffectHandler:
			DefaultEventHandlers.EntryEffect = handler
		case InteractWordHandler:
			DefaultEventHandlers.InteractWord = handler
		case ComboSendHandler:
			DefaultEventHandlers.ComboSend = handler
		case FansChangeHandler:
			DefaultEventHandlers.FansChange = handler
		case InteractiveGameHandler:
			DefaultEventHandlers.InteractiveGame = handler
		case OnlineRankCountHandler:
			DefaultEventHandlers.OnlineRankCount = handler
		case HotRankChangedHandler:
			DefaultEventHandlers.HotRankChanged = handler
		case HotRankChangedV2Handler:
			DefaultEventHandlers.HotRankChangedV2 = handler
		case HotRankSettlementHandler:
			DefaultEventHandlers.HotRankSettlement = handler
		case HotRankSettlementV2Handler:
			DefaultEventHandlers.HotRankSettlementV2 = handler
		case OnlineRankTop3Handler:
			DefaultEventHandlers.OnlineRankTop3 = handler
		case OnlineRankV2Handler:
			DefaultEventHandlers.OnlineRankV2 = handler
		case StopLiveRoomListHandler:
			DefaultEventHandlers.StopLiveRoomList = handler
		case WatchedChangeHandler:
			DefaultEventHandlers.WatchedChange = handler
		case WidgetBannerHandler:
			DefaultEventHandlers.WidgetBanner = handler
		case PopularityRedPocketStartHandler:
			DefaultEventHandlers.PopularityRedPocketStart = handler
		case PopularityRedPocketWinnerListHandler:
			DefaultEventHandlers.PopularityRedPocketWinnerList = handler
		case NoticeMsgHandler:
			DefaultEventHandlers.NoticeMsg = handler
		default:
			log.Errorf("未知handler类型: %T", handler)
			os.Exit(1)
		}
	}
}

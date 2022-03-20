package websocket

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/botplayerneo/bili-live-api/dto"
	"github.com/botplayerneo/bili-live-api/log"
)

type eventPayloadHandler func(*dto.WSPayload)

// 感官上十分冗余，待思考更好的方案
var eventPayloadHandlerMap = map[dto.EventType]eventPayloadHandler{
	dto.EventDanmaku:                       danmakuHandler,
	dto.EventGift:                          giftHandler,
	dto.EventSuperChat:                     superChatHandler,
	dto.EventSuperChatDelete:               superChatDeleteHandler,
	dto.EventGuard:                         guardHandler,
	dto.EventLive:                          liveHandler,
	dto.EventPreparing:                     preparingHandler,
	dto.EventEntryEffect:                   entryEffectHandler,
	dto.EventInteractWord:                  interactWordHandler,
	dto.EventComboSend:                     comboSendHandler,
	dto.EventFansChange:                    fansChangeHandler,
	dto.EventInteractiveGame:               interactiveGameHandler,
	dto.EventOnlineRankCount:               onlineRankCountHandler,
	dto.EventHotRankChanged:                hotRankChangedHandler,
	dto.EventHotRankChangedV2:              hotRankChangedV2Handler,
	dto.EventHotRankSettlement:             hotRankSettlementHandler,
	dto.EventHotRankSettlementV2:           hotRankSettlementV2Handler,
	dto.EventOnlineRankTop3:                onlineRankTop3Handler,
	dto.EventOnlineRankV2:                  onlineRankV2Handler,
	dto.EventStopLiveRoomList:              stopLiveRoomListHandler,
	dto.EventWatchedChange:                 watchedChangeHandler,
	dto.EventWidgetBanner:                  widgetBannerHandler,
	dto.EventPopularityRedPocketStart:      popularityRedPocketStartHandler,
	dto.EventPopularityRedPocketWinnerList: popularityRedPocketWinnerListHandler,
	dto.EventNoticeMsg:                     noticeMsgHandler,
	dto.EventAnchorLotAward:                anchorLotAwardHandler,
	dto.EventUserToastMsg:                  userToastMsgHandler,
	dto.EventRoomChange:                    roomChangeHandler,
	dto.EventRoomBlockMsg:                  roomBlockMsgHandler,
	dto.EventMatchRoomConf:                 matchRoomConfHandler,
	dto.EventCommonNoticeDanmaku:           commonNoticeDanmakuHandler,
	dto.EventAnchorLotCheckStatus:          anchorLotCheckStatusHandler,
	dto.EventAnchorLotEnd:                  anchorLotEndHandler,
	dto.EventAnchorLotStart:                anchorLotStartHandler,
	dto.EventTradingScore:                  tradingScoreHandler,
	// handler_map(for hygen)
}

func danmakuHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Danmaku == nil {
		return
	}
	info := jsoniter.Get(payload.Body, "info")
	d := &dto.Danmaku{
		UID:       info.Get(2, 0).ToInt(),
		Uname:     info.Get(2, 1).ToString(),
		Content:   info.Get(1).ToString(),
		Timestamp: info.Get(0, 4).ToInt64(),
	}
	DefaultEventHandlers.Danmaku(d)
}

func giftHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Gift == nil {
		return
	}
	g := &dto.Gift{}
	jsoniter.Get(payload.Body, "data").ToVal(g)
	DefaultEventHandlers.Gift(g)
}

func superChatHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.SuperChat == nil {
		return
	}
	sc := &dto.SuperChat{}
	jsoniter.Get(payload.Body, "data").ToVal(sc)
	DefaultEventHandlers.SuperChat(sc)
}

func superChatDeleteHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.SuperChatDelete == nil {
		return
	}
	scd := &dto.SuperChatDelete{}
	jsoniter.Get(payload.Body, "data").ToVal(scd)
	DefaultEventHandlers.SuperChatDelete(scd)
}

func guardHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Guard == nil {
		return
	}
	g := &dto.Guard{}
	jsoniter.Get(payload.Body, "data").ToVal(g)
	DefaultEventHandlers.Guard(g)
}

func liveHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Live == nil {
		return
	}
	DefaultEventHandlers.Live()
}

func preparingHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.Preparing == nil {
		return
	}
	DefaultEventHandlers.Preparing()
}

func entryEffectHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.EntryEffect == nil {
		return
	}
	ee := &dto.EntryEffect{}
	jsoniter.Get(payload.Body, "data").ToVal(ee)
	DefaultEventHandlers.EntryEffect(ee)
}

func interactWordHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.InteractWord == nil {
		return
	}
	iw := &dto.InteractWord{}
	jsoniter.Get(payload.Body, "data").ToVal(iw)
	DefaultEventHandlers.InteractWord(iw)
}

func comboSendHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.ComboSend == nil {
		return
	}
	cs := &dto.ComboSend{}
	jsoniter.Get(payload.Body, "data").ToVal(cs)
	DefaultEventHandlers.ComboSend(cs)
}

func fansChangeHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.FansChange == nil {
		return
	}
	fc := &dto.FansChange{}
	jsoniter.Get(payload.Body, "data").ToVal(fc)
	DefaultEventHandlers.FansChange(fc)
}

func interactiveGameHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.InteractiveGame == nil {
		return
	}
	ig := &dto.InteractiveGame{}
	jsoniter.Get(payload.Body, "data").ToVal(ig)
	DefaultEventHandlers.InteractiveGame(ig)
}

func onlineRankCountHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.OnlineRankCount == nil {
		return
	}
	orc := &dto.OnlineRankCount{}
	jsoniter.Get(payload.Body, "data").ToVal(orc)
	DefaultEventHandlers.OnlineRankCount(orc)
}

func hotRankChangedHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.HotRankChanged == nil {
		return
	}
	data := &dto.HotRankChanged{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.HotRankChanged(data)
}

func hotRankChangedV2Handler(payload *dto.WSPayload) {
	if DefaultEventHandlers.HotRankChangedV2 == nil {
		return
	}
	data := &dto.HotRankChangedV2{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.HotRankChangedV2(data)
}

func hotRankSettlementHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.HotRankSettlement == nil {
		return
	}
	data := &dto.HotRankSettlement{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.HotRankSettlement(data)
}

func hotRankSettlementV2Handler(payload *dto.WSPayload) {
	if DefaultEventHandlers.HotRankSettlementV2 == nil {
		return
	}
	data := &dto.HotRankSettlementV2{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.HotRankSettlementV2(data)
}

func onlineRankTop3Handler(payload *dto.WSPayload) {
	if DefaultEventHandlers.OnlineRankTop3 == nil {
		return
	}
	data := &dto.OnlineRankTop3{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.OnlineRankTop3(data)
}

func onlineRankV2Handler(payload *dto.WSPayload) {
	if DefaultEventHandlers.OnlineRankV2 == nil {
		return
	}
	data := &dto.OnlineRankV2{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.OnlineRankV2(data)
}

func stopLiveRoomListHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.StopLiveRoomList == nil {
		return
	}
	data := &dto.StopLiveRoomList{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.StopLiveRoomList(data)
}

func watchedChangeHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.WatchedChange == nil {
		return
	}
	data := &dto.WatchedChange{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.WatchedChange(data)
}

func widgetBannerHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.WidgetBanner == nil {
		return
	}
	data := &dto.WidgetBanner{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.WidgetBanner(data)
}

func popularityRedPocketStartHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.PopularityRedPocketStart == nil {
		return
	}
	data := &dto.PopularityRedPocketStart{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.PopularityRedPocketStart(data)
}

func popularityRedPocketWinnerListHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.PopularityRedPocketWinnerList == nil {
		return
	}
	data := &dto.PopularityRedPocketWinnerList{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.PopularityRedPocketWinnerList(data)
}

func noticeMsgHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.NoticeMsg == nil {
		return
	}
	data := &dto.NoticeMsg{}
	err := jsoniter.Unmarshal(payload.Body, data)
	if err != nil {
		log.Warnf("noticeMsgHandler unmarshal error: %v", err)
		return
	}
	DefaultEventHandlers.NoticeMsg(data)
}

func unknownEventHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.UnknownEvent == nil {
		return
	}
	DefaultEventHandlers.UnknownEvent(payload)
}

func anchorLotAwardHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.AnchorLotAward == nil {
		return
	}
	data := &dto.AnchorLotAward{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.AnchorLotAward(data)
}

func userToastMsgHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.UserToastMsg == nil {
		return
	}
	data := &dto.UserToastMsg{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.UserToastMsg(data)
}

func roomChangeHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.RoomChange == nil {
		return
	}
	data := &dto.RoomChange{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.RoomChange(data)
}

func roomBlockMsgHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.RoomBlockMsg == nil {
		return
	}
	data := &dto.RoomBlockMsg{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.RoomBlockMsg(data)
}

func matchRoomConfHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.MatchRoomConf == nil {
		return
	}
	data := &dto.MatchRoomConf{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.MatchRoomConf(data)
}

func commonNoticeDanmakuHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.CommonNoticeDanmaku == nil {
		return
	}
	data := &dto.CommonNoticeDanmaku{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.CommonNoticeDanmaku(data)
}

func anchorLotCheckStatusHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.AnchorLotCheckstatus == nil {
		return
	}
	data := &dto.AnchorLotCheckStatus{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.AnchorLotCheckstatus(data)
}

func anchorLotEndHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.AnchorLotEnd == nil {
		return
	}
	data := &dto.AnchorLotEnd{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.AnchorLotEnd(data)
}

func anchorLotStartHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.AnchorLotStart == nil {
		return
	}
	data := &dto.AnchorLotStart{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.AnchorLotStart(data)
}

func tradingScoreHandler(payload *dto.WSPayload) {
	if DefaultEventHandlers.TradingScore == nil {
		return
	}
	data := &dto.TradingScore{}
	jsoniter.Get(payload.Body, "data").ToVal(data)
	DefaultEventHandlers.TradingScore(data)
}

// handler_impl(for hygen)

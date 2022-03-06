package dto

type EventType string

// cmd事件类型
const (
	EventDanmaku                       EventType = "DANMU_MSG"
	EventGift                          EventType = "SEND_GIFT"
	EventSuperChat                     EventType = "SUPER_CHAT_MESSAGE"
	EventSuperChatDelete               EventType = "SUPER_CHAT_MESSAGE_DELETE"
	EventGuard                         EventType = "GUARD_BUY"
	EventLive                          EventType = "LIVE"
	EventPreparing                     EventType = "PREPARING"
	EventEntryEffect                   EventType = "ENTRY_EFFECT"
	EventInteractWord                  EventType = "INTERACT_WORD"
	EventComboSend                     EventType = "COMBO_SEND"
	EventFansChange                    EventType = "ROOM_REAL_TIME_MESSAGE_UPDATE"
	EventInteractiveGame               EventType = "LIVE_INTERACTIVE_GAME"
	EventOnlineRankCount               EventType = "ONLINE_RANK_COUNT"
	EventHotRankChanged                EventType = "HOT_RANK_CHANGED"
	EventHotRankChangedV2              EventType = "HOT_RANK_CHANGED_V2"
	EventHotRankSettlement             EventType = "HOT_RANK_SETTLEMENT"
	EventHotRankSettlementV2           EventType = "HOT_RANK_SETTLEMENT_V2"
	EventOnlineRankTop3                EventType = "ONLINE_RANK_TOP3"
	EventOnlineRankV2                  EventType = "ONLINE_RANK_V2"
	EventStopLiveRoomList              EventType = "STOP_LIVE_ROOM_LIST"
	EventWatchedChange                 EventType = "WATCHED_CHANGE"
	EventWidgetBanner                  EventType = "WIDGET_BANNER"
	EventPopularityRedPocketStart      EventType = "POPULARITY_RED_POCKET_START"
	EventPopularityRedPocketWinnerList EventType = "POPULARITY_RED_POCKET_WINNER_LIST"
	EventNoticeMsg                     EventType = "NOTICE_MSG"
)

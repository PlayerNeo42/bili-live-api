package main

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/robfig/cron"
	"github.com/spellingDragon/bili-live-api/dto"
	"github.com/spellingDragon/bili-live-api/log"
	"github.com/spellingDragon/bili-live-api/websocket"
	"time"
)

func main() {
	jsonStr := "{\"cmd\":\"SUPER_CHAT_MESSAGE\",\"data\":{\"background_bottom_color\":\"#2A60B2\",\"background_color\":\"#EDF5FF\",\"background_color_end\":\"#405D85\",\"background_color_start\":\"#3171D2\",\"background_icon\":\"\",\"background_image\":\"https://i0.hdslb.com/bfs/live/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png\",\"background_price_color\":\"#7497CD\",\"color_point\":0.7,\"dmscore\":72,\"end_time\":1656952477,\"gift\":{\"gift_id\":12000,\"gift_name\":\"醒目留言\",\"num\":1},\"id\":4446245,\"is_ranked\":1,\"is_send_audit\":0,\"medal_info\":{\"anchor_roomid\":22676119,\"anchor_uname\":\"猫雷NyaRu_Official\",\"guard_level\":0,\"icon_id\":0,\"is_lighted\":1,\"medal_color\":\"#1a544b\",\"medal_color_border\":1725515,\"medal_color_end\":5414290,\"medal_color_start\":1725515,\"medal_level\":24,\"medal_name\":\"喵喵露\",\"special\":\"\",\"target_id\":697091119},\"message\":\"今天不管干啥，能一直开着AI字幕吗？\",\"message_font_color\":\"#A3F6FF\",\"message_trans\":\"\",\"price\":30,\"rate\":1000,\"start_time\":1656952417,\"time\":60,\"token\":\"EE957DAB\",\"trans_mark\":0,\"ts\":1656952417,\"uid\":32629114,\"user_info\":{\"face\":\"http://i1.hdslb.com/bfs/face/f21caae6ee4b68e47421120b74627bc5262e52ca.jpg\",\"face_frame\":\"\",\"guard_level\":0,\"is_main_vip\":1,\"is_svip\":0,\"is_vip\":0,\"level_color\":\"#61c05a\",\"manager\":0,\"name_color\":\"#666666\",\"title\":\"0\",\"uname\":\"羽游仙\",\"user_level\":18}},\"roomid\":22676119}"
	sc := &dto.SuperChat{}
	println(jsonStr)
	data := jsoniter.Get([]byte(jsonStr), "data")
	println(data.ToString())
	data.ToVal(sc)
	result, _ := json.Marshal(sc)
	fmt.Printf("%s\n", result)
	timeStamp := time.Unix(1656952417, 0)
	println(timeStamp.Format("2006-01-02 03:04:05"))
	timer := getTimer()
	timer.Start()
	c := make(chan bool)
	<-c
}

func getTimer() *cron.Cron {
	timer := cron.New()
	timer.AddFunc("0 0/1 * * * ?", func() {
		// TODO 失败？
		log.Infof("测试定时器")
	})
	return timer
}

func danmakuHandler() websocket.DanmakuHandler {
	return func(danmaku *dto.Danmaku) {
		fmt.Printf("%s:%s\n", danmaku.Uname, danmaku.Content)
	}
}

func giftHandler() websocket.GiftHandler {
	return func(gift *dto.Gift) {
		fmt.Printf("%s 赠送 [礼物%.1f￥]%s x %d\n", gift.Uname, float64(gift.Price)/1000, gift.GiftName, gift.Num)
	}
}

func guardHandler() websocket.GuardHandler {
	return func(guard *dto.Guard) {
		fmt.Printf("%s 开通舰长,级别%d", guard.Username, guard.GuardLevel)
	}
}

func superChatHandler() websocket.SuperChatHandler {
	return func(superChat *dto.SuperChat) {
		fmt.Printf("[%d￥SC]%s:%s\n", superChat.Price, superChat.UserInfo.Uname, superChat.Message)
	}
}

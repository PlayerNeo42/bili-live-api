package resource

import (
	"fmt"
)

// Dynamic 动态
type Dynamic struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Message string `json:"message"`
	Data    struct {
		HasMore int `json:"has_more"`
		Cards   []struct {
			Desc struct {
				Uid         int   `json:"uid"`
				Type        int   `json:"type"`
				Rid         int64 `json:"rid"`
				View        int   `json:"view"`
				Repost      int   `json:"repost"`
				Comment     int   `json:"comment"`
				Like        int   `json:"like"`
				IsLiked     int   `json:"is_liked"`
				DynamicId   int64 `json:"dynamic_id"`
				Timestamp   int64 `json:"timestamp"`
				OrigType    int   `json:"orig_type"`
				UserProfile struct {
					Info struct {
						Uid     int    `json:"uid"`
						Uname   string `json:"uname"`
						Face    string `json:"face"`
						FaceNft int    `json:"face_nft"`
					} `json:"info"`
					Card struct {
						OfficialVerify struct {
							Type int    `json:"type"`
							Desc string `json:"desc"`
						} `json:"official_verify"`
					} `json:"card"`
					Vip struct {
						VipType    int   `json:"vipType"`
						VipDueDate int64 `json:"vipDueDate"`
						VipStatus  int   `json:"vipStatus"`
						ThemeType  int   `json:"themeType"`
						Label      struct {
							Path        string `json:"path"`
							Text        string `json:"text"`
							LabelTheme  string `json:"label_theme"`
							TextColor   string `json:"text_color"`
							BgStyle     int    `json:"bg_style"`
							BgColor     string `json:"bg_color"`
							BorderColor string `json:"border_color"`
						} `json:"label"`
						AvatarSubscript    int    `json:"avatar_subscript"`
						NicknameColor      string `json:"nickname_color"`
						Role               int    `json:"role"`
						AvatarSubscriptUrl string `json:"avatar_subscript_url"`
					} `json:"vip"`
					Pendant struct {
						Pid               int    `json:"pid"`
						Name              string `json:"name"`
						Image             string `json:"image"`
						Expire            int    `json:"expire"`
						ImageEnhance      string `json:"image_enhance"`
						ImageEnhanceFrame string `json:"image_enhance_frame"`
					} `json:"pendant"`
					DecorateCard struct {
						Mid          int    `json:"mid"`
						Id           int    `json:"id"`
						CardUrl      string `json:"card_url"`
						CardType     int    `json:"card_type"`
						Name         string `json:"name"`
						ExpireTime   int    `json:"expire_time"`
						CardTypeName string `json:"card_type_name"`
						Uid          int    `json:"uid"`
						ItemId       int    `json:"item_id"`
						ItemType     int    `json:"item_type"`
						BigCardUrl   string `json:"big_card_url"`
						JumpUrl      string `json:"jump_url"`
						Fan          struct {
							IsFan   int    `json:"is_fan"`
							Number  int    `json:"number"`
							Color   string `json:"color"`
							NumDesc string `json:"num_desc"`
						} `json:"fan"`
						ImageEnhance string `json:"image_enhance"`
					} `json:"decorate_card"`
					Rank      string `json:"rank"`
					Sign      string `json:"sign"`
					LevelInfo struct {
						CurrentLevel int `json:"current_level"`
					} `json:"level_info"`
				} `json:"user_profile"`
				UidType      int    `json:"uid_type"`
				Status       int    `json:"status"`
				DynamicIdStr string `json:"dynamic_id_str"`
				PreDyIdStr   string `json:"pre_dy_id_str"`
				OrigDyIdStr  string `json:"orig_dy_id_str"`
				RidStr       string `json:"rid_str"`
				Acl          int    `json:"acl,omitempty"`
				PreDyId      int64  `json:"pre_dy_id,omitempty"`
				OrigDyId     int64  `json:"orig_dy_id,omitempty"`
				Stype        int    `json:"stype,omitempty"`
				RType        int    `json:"r_type,omitempty"`
				InnerId      int    `json:"inner_id,omitempty"`
				Origin       struct {
					Uid          int    `json:"uid"`
					Type         int    `json:"type"`
					Rid          int    `json:"rid"`
					Acl          int    `json:"acl"`
					View         int    `json:"view"`
					Repost       int    `json:"repost"`
					Like         int    `json:"like"`
					DynamicId    int64  `json:"dynamic_id"`
					Timestamp    int    `json:"timestamp"`
					PreDyId      int    `json:"pre_dy_id"`
					OrigDyId     int    `json:"orig_dy_id"`
					UidType      int    `json:"uid_type"`
					Stype        int    `json:"stype"`
					RType        int    `json:"r_type"`
					InnerId      int    `json:"inner_id"`
					Status       int    `json:"status"`
					DynamicIdStr string `json:"dynamic_id_str"`
					PreDyIdStr   string `json:"pre_dy_id_str"`
					OrigDyIdStr  string `json:"orig_dy_id_str"`
					RidStr       string `json:"rid_str"`
					Bvid         string `json:"bvid"`
				} `json:"origin,omitempty"`
			} `json:"desc"`
			Card       string `json:"card"`
			ExtendJson string `json:"extend_json"`
			Extra      struct {
				IsSpaceTop int `json:"is_space_top"`
			} `json:"extra"`
			Display struct {
				EmojiInfo struct {
					EmojiDetails []struct {
						EmojiName string `json:"emoji_name"`
						Id        int    `json:"id"`
						PackageId int    `json:"package_id"`
						State     int    `json:"state"`
						Type      int    `json:"type"`
						Attr      int    `json:"attr"`
						Text      string `json:"text"`
						Url       string `json:"url"`
						Meta      struct {
							Size int `json:"size"`
						} `json:"meta"`
						Mtime int `json:"mtime"`
					} `json:"emoji_details"`
				} `json:"emoji_info,omitempty"`
				Relation struct {
					Status     int `json:"status"`
					IsFollow   int `json:"is_follow"`
					IsFollowed int `json:"is_followed"`
				} `json:"relation"`
				CommentInfo struct {
					CommentIds string `json:"comment_ids"`
				} `json:"comment_info"`
				Origin struct {
					TopicInfo struct {
						TopicDetails []struct {
							TopicId    int    `json:"topic_id"`
							TopicName  string `json:"topic_name"`
							IsActivity int    `json:"is_activity"`
							TopicLink  string `json:"topic_link"`
						} `json:"topic_details"`
					} `json:"topic_info"`
					UsrActionTxt string `json:"usr_action_txt"`
					Relation     struct {
						Status     int `json:"status"`
						IsFollow   int `json:"is_follow"`
						IsFollowed int `json:"is_followed"`
					} `json:"relation"`
					ShowTip struct {
						DelTip string `json:"del_tip"`
					} `json:"show_tip"`
					CoverPlayIconUrl string `json:"cover_play_icon_url"`
				} `json:"origin,omitempty"`
			} `json:"display"`
		} `json:"cards"`
		NextOffset int64 `json:"next_offset"`
		Gt         int   `json:"_gt_"`
	} `json:"data"`
}

func GetDynamic(uid int64) (*Dynamic, error) {
	result := &Dynamic{}
	_, err := vcApiClient.R().
		SetQueryParam("host_uid", fmt.Sprintf("%d", uid)).
		SetResult(result).
		Get("/dynamic_svr/v1/dynamic_svr/space_history")
	return result, err
}

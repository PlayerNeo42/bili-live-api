package dto

type EntryEffect struct {
	Id               int           `json:"id"`
	Uid              int           `json:"uid"`
	TargetId         int           `json:"target_id"`
	MockEffect       int           `json:"mock_effect"`
	Face             string        `json:"face"`
	PrivilegeType    int           `json:"privilege_type"`
	CopyWriting      string        `json:"copy_writing"`
	CopyColor        string        `json:"copy_color"`
	HighlightColor   string        `json:"highlight_color"`
	Priority         int           `json:"priority"`
	BasemapUrl       string        `json:"basemap_url"`
	ShowAvatar       int           `json:"show_avatar"`
	EffectiveTime    int           `json:"effective_time"`
	WebBasemapUrl    string        `json:"web_basemap_url"`
	WebEffectiveTime int           `json:"web_effective_time"`
	WebEffectClose   int           `json:"web_effect_close"`
	WebCloseTime     int           `json:"web_close_time"`
	Business         int           `json:"business"`
	CopyWritingV2    string        `json:"copy_writing_v2"`
	IconList         []interface{} `json:"icon_list"`
	MaxDelayTime     int           `json:"max_delay_time"`
	TriggerTime      int64         `json:"trigger_time"`
	Identities       int           `json:"identities"`
}

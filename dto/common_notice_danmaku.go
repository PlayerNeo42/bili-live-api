package dto

// CommonNoticeDanmaku 提示信息,如'恭喜主播完成"小小花束"任务'
type CommonNoticeDanmaku struct {
	ContentSegments []struct {
		FontColor string `json:"font_color"`
		Text      string `json:"text"`
		Type      int    `json:"type"`
	} `json:"content_segments"`
	DmScore   int   `json:"dmscore"`
	Terminals []int `json:"terminals"`
}

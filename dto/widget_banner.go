package dto

type WidgetBanner struct {
	Timestamp  int `json:"timestamp"`
	WidgetList map[string]struct {
		ID             int      `json:"id"`
		Title          string   `json:"title"`
		Cover          string   `json:"cover"`
		WebCover       string   `json:"web_cover"`
		TipText        string   `json:"tip_text"`
		TipTextColor   string   `json:"tip_text_color"`
		TipBottomColor string   `json:"tip_bottom_color"`
		JumpURL        string   `json:"jump_url"`
		URL            string   `json:"url"`
		StayTime       int      `json:"stay_time"`
		Site           int      `json:"site"`
		PlatformIn     []string `json:"platform_in"`
		Type           int      `json:"type"`
		BandID         int      `json:"band_id"`
		SubKey         string   `json:"sub_key"`
		SubData        string   `json:"sub_data"`
		IsAdd          bool     `json:"is_add"`
	} `json:"widget_list"`
}

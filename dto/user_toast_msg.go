package dto

// UserToastMsg 弹幕栏中显示的toast信息，如xxx自动续费了舰长，xxx开通了舰长
type UserToastMsg struct {
	AnchorShow       bool   `json:"anchor_show"`
	Color            string `json:"color"`
	DmScore          int    `json:"dmscore"`
	EffectID         int    `json:"effect_id"`
	EndTime          int    `json:"end_time"`
	GuardLevel       int    `json:"guard_level"`
	IsShow           int    `json:"is_show"`
	Num              int    `json:"num"`
	OpType           int    `json:"op_type"`
	PayflowID        string `json:"payflow_id"`
	Price            int    `json:"price"`
	RoleName         string `json:"role_name"`
	StartTime        int    `json:"start_time"`
	SvgaBlock        int    `json:"svga_block"`
	TargetGuardCount int    `json:"target_guard_count"`
	ToastMsg         string `json:"toast_msg"`
	UID              int    `json:"uid"`
	Unit             string `json:"unit"`
	UserShow         bool   `json:"user_show"`
	Username         string `json:"username"`
}

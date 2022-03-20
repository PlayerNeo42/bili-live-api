package dto

// Danmaku 转换后的弹幕对象，并非原始数据
// TODO 完善该类型
type Danmaku struct {
	UID       int
	Uname     string
	Content   string
	Timestamp int64
}

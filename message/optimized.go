package message

type Danmaku struct {
	Uid       int
	Uname     string
	Content   string
	Timestamp int
}

type SuperChat struct {
	Uid       int
	Uname     string
	Message   string
	Price     int // 单位 1000/元
}

type Gift struct {
	Uid       int
	Uname     string
	Price     int // 单位 1000/元
	GiftName  string
}

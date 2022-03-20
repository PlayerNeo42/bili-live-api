package dto

// AnchorLotCheckStatus 天选之人相关
type AnchorLotCheckStatus struct {
	ID           int    `json:"id"`
	RejectReason string `json:"reject_reason"`
	Status       int    `json:"status"`
	UID          int    `json:"uid"`
}

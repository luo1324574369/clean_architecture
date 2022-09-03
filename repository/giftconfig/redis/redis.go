package redis

// PO 非必须实现,可以直接使用entity代替
type PO struct {
	ID         uint64 `json:"id"`
	GiftList   string `json:"gift_list"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
}

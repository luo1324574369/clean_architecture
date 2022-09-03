package mysql

type PO struct {
	ID         uint64 `orm:"id"`
	GiftList   string `orm:"gift_list"`
	CreateTime int64  `orm:"create_time"`
	ModifyTime int64  `orm:"modify_time"`
}

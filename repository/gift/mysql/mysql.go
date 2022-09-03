package mysql

type PO struct {
	ID         uint64 `orm:"id"`
	UIN        uint64 `orm:"uin"`
	ConfigID   uint64 `orm:"config_id"`
	CreateTime int64  `orm:"create_time"`
	ModifyTime int64  `orm:"modify_time"`
}

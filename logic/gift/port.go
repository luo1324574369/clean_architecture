package gift

import (
	"context"

	gift "github.com/luo1324574369/clean_architecture/enity/gift"
)

type AddGiftConfigCMD struct {
}

type GetConfigListRQY struct {
}

type Port interface {
	AddGiftConfig(ctx context.Context, cmd AddGiftConfigCMD) error

	GetGiftConfigList(ctx context.Context, qry GetConfigListRQY) []*gift.ConfigEntity

	AddGift(ctx context.Context, uin uint64, ConfigID uint64) error
}

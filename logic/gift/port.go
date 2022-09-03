package gift

import (
	"context"
)

type AddGiftConfigCMD struct {
}

type GetConfigListRQY struct {
}

type Port interface {
	AddGiftConfig(ctx context.Context, cmd AddGiftConfigCMD) error

	AddGift(ctx context.Context, uin uint64) error
}

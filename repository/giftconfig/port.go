package giftconfig

import (
	"context"

	"github.com/luo1324574369/clean_architecture/enity/gift"
)

type GiftConfigQRY struct {
}

type GiftConfigPort interface {
	Create(ctx context.Context, entity *gift.ConfigEntity) error
	List(ctx context.Context, qry GiftConfigQRY) []*gift.ConfigEntity
	RefreshCache(ctx context.Context) error
}

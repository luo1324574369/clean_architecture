package giftconfig

import (
	"context"

	"github.com/luo1324574369/clean_architecture/enity/gift"
)

type Adapter struct {
}

func (a *Adapter) Create(ctx context.Context, entity *gift.ConfigEntity) error {
	return nil
}

func (a *Adapter) List(ctx context.Context,
	qry GiftConfigQRY) []*gift.ConfigEntity {
	return nil
}

func (a *Adapter) RefreshCache(ctx context.Context) error {
	return nil
}

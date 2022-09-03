package gift

import (
	"context"
	"errors"
	"time"

	"github.com/luo1324574369/clean_architecture/enity/async"
	giftEntity "github.com/luo1324574369/clean_architecture/enity/gift"
	"github.com/luo1324574369/clean_architecture/enity/log"
	"github.com/luo1324574369/clean_architecture/repository/gift"
	"github.com/luo1324574369/clean_architecture/repository/giftconfig"
)

type Adapter struct {
	giftPort       gift.GiftPort
	giftConfigPort giftconfig.GiftConfigPort
}

func NewAdapter() *Adapter {
	return &Adapter{
		giftPort:       &gift.Adapter{},
		giftConfigPort: &giftconfig.Adapter{},
	}
}

func (a *Adapter) AddGiftConfig(ctx context.Context, cmd AddGiftConfigCMD) error {
	entity := &giftEntity.ConfigEntity{}
	if err := a.giftConfigPort.Create(ctx, entity); err != nil {
		log.ErrorContextf(ctx, "xxx error:", err)
		return err
	}

	_ = async.Go(ctx, 3*time.Second, func(cloneCtx context.Context) {
		_ = a.giftConfigPort.RefreshCache(cloneCtx)
	})

	return nil
}

func (a *Adapter) GetGiftConfigList(ctx context.Context,
	qry GetConfigListRQY) []*giftEntity.ConfigEntity {
	q := giftconfig.GiftConfigQRY{}
	return a.giftConfigPort.List(ctx, q)
}

func (a *Adapter) AddGift(ctx context.Context, uin uint64, ConfigID uint64) error {
	configs := a.giftConfigPort.List(ctx, giftconfig.GiftConfigQRY{})

	var config *giftEntity.ConfigEntity
	for _, c := range configs {
		if c.ID == ConfigID {
			config = c
			break
		}
	}

	if config == nil {
		return errors.New("not valid config_id")
	}

	if config.HasQB() {
		// ... 检查是否绑定qq
	}

	insertData := &giftEntity.Entity{}
	if err := a.giftPort.Create(ctx, insertData); err != nil {
		return err
	}
	return nil
}

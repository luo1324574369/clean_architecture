package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luo1324574369/clean_architecture/logic/gift"
)

type giftSvr struct {
	logic gift.Port
}

func NewGiftServer() *giftSvr {
	return &giftSvr{
		logic: gift.NewAdapter(),
	}
}

func (p *giftSvr) AddGiftConfig(ctx *gin.Context) {
	cmd := gift.AddGiftConfigCMD{}
	if err := p.logic.AddGiftConfig(ctx, cmd); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	}
}

func (p *giftSvr) GetGiftConfigList(ctx *gin.Context) {
	entityList := p.logic.GetGiftConfigList(ctx, gift.GetConfigListRQY{})
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": entityList,
	})
}

func (p *giftSvr) AddGift(ctx *gin.Context) {
	// ...
}

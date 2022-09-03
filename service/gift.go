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

// AddGiftConfig 增加配置
func (p *giftSvr) AddGiftConfig(ctx *gin.Context) {
	cmd := gift.AddGiftConfigCMD{}
	if err := p.logic.AddGiftConfig(ctx, cmd); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "fail",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// AddGift 判断某个人是否满足条件,送礼物
func (p *giftSvr) AddGift(ctx *gin.Context) {
	if err := p.logic.AddGift(ctx, 1); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "fail",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

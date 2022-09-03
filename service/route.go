package service

import (
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func init() {
	Server = gin.Default()
	InitRoute()
}

func InitRoute() {
	giftSvr := NewGiftServer()

	Server.GET("/addGift", giftSvr.AddGift)
	Server.GET("/addGiftConfig", giftSvr.AddGiftConfig)
	Server.GET("/GetGiftConfigList", giftSvr.GetGiftConfigList)
}

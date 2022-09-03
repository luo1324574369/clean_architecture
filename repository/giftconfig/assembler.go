package giftconfig

import (
	"encoding/json"

	"github.com/luo1324574369/clean_architecture/repository/giftconfig/redis"

	"github.com/luo1324574369/clean_architecture/enity/gift"
	"github.com/luo1324574369/clean_architecture/enity/log"
	"github.com/luo1324574369/clean_architecture/repository/giftconfig/mysql"
)

func entityToMysqlPO(entity *gift.ConfigEntity) *mysql.PO {
	giftListJson, err := json.Marshal(entity)
	if err != nil {
		giftListJson = []byte("")
		log.Errorf("error %v", err)
	}

	return &mysql.PO{
		ID:       entity.ID,
		GiftList: string(giftListJson),
	}
}

func mysqlToRedisPO(mysqlPO *mysql.PO) *redis.PO {
	return &redis.PO{
		ID:         mysqlPO.ID,
		GiftList:   mysqlPO.GiftList,
		CreateTime: mysqlPO.CreateTime,
		ModifyTime: mysqlPO.ModifyTime,
	}
}

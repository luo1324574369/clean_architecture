package gift

import (
	"github.com/luo1324574369/clean_architecture/enity/gift"
	"github.com/luo1324574369/clean_architecture/repository/gift/mysql"
)

func entityToMysqlPO(entity *gift.Entity) *mysql.PO {
	return &mysql.PO{
		ID:       entity.ID,
		UIN:      entity.UIN,
		ConfigID: entity.ConfigID,
	}
}

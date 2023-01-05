package db

import (
	"encoding/json"
	"github.com/cdgProcessor/outboundWriter/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Writer(ch <-chan []byte) {
	zap.L().Info("Inbound DB connection start")
	dsn := "root:Welcome@1@tcp(172.25.240.10:30306)/cdg?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Panic("open mysql failed, err: ", zap.Error(err))
	}

	// 初始化数据库
	db.AutoMigrate(&models.MbRc{})

	var message models.MbRc

	for {
		sms := <-ch
		json.Unmarshal(sms, &message)
		db.Create(&message)
	}
}

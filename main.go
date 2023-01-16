package main

import (
	"github.com/cdgProcessor/outboundWriter/db"
	"github.com/cdgProcessor/outboundWriter/logger"
	"github.com/cdgProcessor/outboundWriter/messageQ"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger("./logs/outboundWriter.log")
	zap.L().Info("Outbound sms DB processor starting...")

	out2dbChan := make(chan []byte)

	go messageQ.MQRead(out2dbChan, "outboundSMS", "outboundSmsToDB", "outboundToDB")

	db.Writer(out2dbChan)
}

package logging

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	var err error
	log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

package internal

import (
	"fmt"

	"api_server/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

func (w *writer) Printf(message string, data ...interface{}) {
	global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
}

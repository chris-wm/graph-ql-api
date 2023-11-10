package model

import (
	"github.com/electivetechnology/utility-library-go/logger"
)

var log logger.Logging

func init() {
	// Add generic logger
	log = logger.NewLogger("model")
}

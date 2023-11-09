package main

import (
	logging "github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/router"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/controllers/status"
)

func main() {
	// Start Main logger
	logger := logging.NewLogger("main")
	logger.DebugF("Starting Graph-ql-api App")

	// Create and start router
	logger.DebugF("Creating new  Router")
	router := router.NewRouter()
	router.Run()
}

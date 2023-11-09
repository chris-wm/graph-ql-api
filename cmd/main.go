package main

import (
	logging "github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/router"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/controllers/status"
)

func main() {
	// Start Main logger
	logger := logging.NewLogger("main")
	logger.DebugF("Starting Graph-ql-api App")

	// Run App Migrations
	adapter := adapter.NewAdapter()
	adapter.Connect()
	adapter.MigrationsMigrate()

	// Create and start router
	logger.DebugF("Creating new  Router")
	router := router.NewRouter()
	router.Run()
}

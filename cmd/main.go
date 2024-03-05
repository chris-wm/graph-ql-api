package main

import (
	logging "github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/router"
	"gitlab.et-ns.net/connect/graph-ql-api/internal/adapter"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/controllers/status"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/controllers/templates"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/controllers/users"
	_ "gitlab.et-ns.net/connect/graph-ql-api/internal/migrations"
)

func main() {
	// Start Main logger
	logger := logging.NewLogger("main")
	logger.DebugF("Starting graph-ql-api App")

	// Run App Migrations
	adapter := adapter.NewAdapter()
	adapter.Connect()
	adapter.MigrationsMigrate()

	// Create and start router
	logger.DebugF("Creating new  Router")
	router := router.NewRouter()
	router.Run()
}

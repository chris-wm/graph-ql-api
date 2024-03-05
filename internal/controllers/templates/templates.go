package templates

import (
	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/router"
)

var log logger.ContextLogging

const (
	ACL_SUBJECT = "TEMPLATE"
)

func init() {
	// Add generic logger
	log = logger.NewLogger("controller/templates")

	// Register route with router
	for _, route := range GetRoutes() {
		router.RegisterRoute(route)
	}
}

func GetRoutes() []router.Route {
	var routes []router.Route

	// OPTIONS /v2/templates
	templateCreateOptions := router.Route{}
	templateCreateOptions.Method = []string{"OPTIONS"}
	templateCreateOptions.Path = "/v2/templates"
	templateCreateOptions.Handler = router.NoContent
	routes = append(routes, templateCreateOptions)

	// OPTIONS /v2/templates/:id
	templateReadOptions := router.Route{}
	templateReadOptions.Method = []string{"OPTIONS"}
	templateReadOptions.Path = "/v2/templates/:id"
	templateReadOptions.Handler = router.NoContent
	routes = append(routes, templateReadOptions)

	// HEAD /v2/templates
	templateCreateHead := router.Route{}
	templateCreateHead.Method = []string{"HEAD"}
	templateCreateHead.Path = "/v2/templates"
	templateCreateHead.Handler = router.NoContent
	templateCreateHead.IsAuthenticated = true
	routes = append(routes, templateCreateHead)

	// HEAD /v2/templates/:id
	templateReadHead := router.Route{}
	templateReadHead.Method = []string{"HEAD"}
	templateReadHead.Path = "/v2/templates/:id"
	templateReadHead.Handler = router.NoContent
	templateReadHead.IsAuthenticated = true
	routes = append(routes, templateReadHead)

	// GET /v2/templates
	templateCreate := router.Route{}
	templateCreate.Method = []string{"POST"}
	templateCreate.Path = "/v2/templates"
	templateCreate.Handler = Create
	routes = append(routes, templateCreate)

	// GET /v2/templates/{id}
	templateRead := router.Route{}
	templateRead.Method = []string{"GET"}
	templateRead.Path = "/v2/templates/:id"
	templateRead.Handler = Read
	routes = append(routes, templateRead)

	// PATCH /v2/templates/{id}
	templateUpdate := router.Route{}
	templateUpdate.Method = []string{"PATCH"}
	templateUpdate.Path = "/v2/templates/:id"
	templateUpdate.Handler = Update
	routes = append(routes, templateUpdate)

	// DELETE /v2/templates/{id}
	templateDelete := router.Route{}
	templateDelete.Method = []string{"DELETE"}
	templateDelete.Path = "/v2/templates/:id"
	templateDelete.Handler = Delete
	routes = append(routes, templateDelete)

	// GET /v2/templates
	templateList := router.Route{}
	templateList.Method = []string{"GET"}
	templateList.Path = "/v2/templates"
	templateList.Handler = List
	routes = append(routes, templateList)

	return routes
}

package users

import (
	"github.com/electivetechnology/utility-library-go/logger"
	"github.com/electivetechnology/utility-library-go/router"
)

var log logger.ContextLogging

const (
	ACL_SUBJECT = "USER"
)

func init() {
	// Add generic logger
	log = logger.NewLogger("controller/users")

	// Register route with router
	for _, route := range GetRoutes() {
		router.RegisterRoute(route)
	}
}

func GetRoutes() []router.Route {
	var routes []router.Route

	// OPTIONS /users
	userCreateOptions := router.Route{}
	userCreateOptions.Method = []string{"OPTIONS"}
	userCreateOptions.Path = "/users"
	userCreateOptions.Handler = router.NoContent
	routes = append(routes, userCreateOptions)

	// OPTIONS /users/:id
	userReadOptions := router.Route{}
	userReadOptions.Method = []string{"OPTIONS"}
	userReadOptions.Path = "/users/:id"
	userReadOptions.Handler = router.NoContent
	routes = append(routes, userReadOptions)

	// HEAD /users
	userCreateHead := router.Route{}
	userCreateHead.Method = []string{"HEAD"}
	userCreateHead.Path = "/users"
	userCreateHead.Handler = router.NoContent
	userCreateHead.IsAuthenticated = true
	routes = append(routes, userCreateHead)

	// HEAD /users/:id
	userReadHead := router.Route{}
	userReadHead.Method = []string{"HEAD"}
	userReadHead.Path = "/users/:id"
	userReadHead.Handler = router.NoContent
	userReadHead.IsAuthenticated = true
	routes = append(routes, userReadHead)

	// GET /users
	userCreate := router.Route{}
	userCreate.Method = []string{"POST"}
	userCreate.Path = "/users"
	userCreate.Handler = Create
	userCreate.IsAuthenticated = true
	routes = append(routes, userCreate)

	// GET /users/{id}
	userRead := router.Route{}
	userRead.Method = []string{"GET"}
	userRead.Path = "/users/:id"
	userRead.Handler = Read
	userRead.IsAuthenticated = true
	routes = append(routes, userRead)

	// PATCH /users/{id}
	userUpdate := router.Route{}
	userUpdate.Method = []string{"PATCH"}
	userUpdate.Path = "/users/:id"
	userUpdate.Handler = Update
	userUpdate.IsAuthenticated = true
	routes = append(routes, userUpdate)

	// DELETE /users/{id}
	userDelete := router.Route{}
	userDelete.Method = []string{"DELETE"}
	userDelete.Path = "/users/:id"
	userDelete.Handler = Delete
	userDelete.IsAuthenticated = true
	routes = append(routes, userDelete)

	// GET /users
	userList := router.Route{}
	userList.Method = []string{"GET"}
	userList.Path = "/users"
	userList.Handler = List
	userList.IsAuthenticated = true
	routes = append(routes, userList)

	return routes
}

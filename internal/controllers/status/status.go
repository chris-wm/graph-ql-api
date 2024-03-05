package status

import (
	"github.com/electivetechnology/utility-library-go/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StatusResponse struct {
	Status int `json:"status"`
}

func init() {
	// Register route with router
	for _, route := range GetRoutes() {
		router.RegisterRoute(route)
	}
}

func GetRoutes() []router.Route {
	var routes []router.Route

	// GET /status
	statusGet := router.Route{}
	statusGet.Method = []string{"GET"}
	statusGet.Path = "/status"
	statusGet.Handler = Get
	routes = append(routes, statusGet)

	// HEAD /v2/status
	statusHead := router.Route{}
	statusHead.Method = []string{"HEAD"}
	statusHead.Path = "/v2/status"
	statusHead.Handler = router.NoContent
	routes = append(routes, statusHead)

	// GET /v2/status
	statusV2 := router.Route{}
	statusV2.Method = []string{"GET"}
	statusV2.Path = "/v2/status"
	statusV2.Handler = Get
	routes = append(routes, statusV2)

	return routes
}

func Get(c *gin.Context) {
	ret := GetData()
	c.JSON(http.StatusOK, ret)
}

func GetData() StatusResponse {
	ret := StatusResponse{}
	ret.Status = http.StatusOK

	return ret
}

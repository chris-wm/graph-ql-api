package users

import (
	"net/http"

	context "github.com/electivetechnology/utility-library-go/request/context"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// Reserve space for errors
	var errors []error

	// Get Limit
	limit, err := context.GetLimit(c)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got limit for request: %v", limit)

	// Get Offset
	offset, err := context.GetOffset(c)
	if err != nil {
		errors = append(errors, err)
	}
	log.Printf("Got offset for request: %v", offset)

	// Output errors (if any) and terminate
	if len(errors) > 0 {
		var list []string
		for _, e := range errors {
			list = append(list, e.Error())
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": list})
		return
	}

	// Get users
	users, err := ListUser(
		limit.GetLimit(),
		offset.GetOffset())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

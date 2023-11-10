package users

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	id := c.Param("id")

	userId, err := uuid.Parse(id)

	if err != nil {
		log.Fatalf("Could not parse %s into uuid", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ReadUser(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

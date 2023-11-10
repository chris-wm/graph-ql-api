package users

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Email    string `json:"email" validate:"omitempty"`
	Password string `json:"password" validate:"omitempty"`
	Name     string `json:"name" validate:"omitempty"`
}

func Update(c *gin.Context) {
	id := c.Param("id")

	// Validate request
	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := uuid.Parse(id)

	if err != nil {
		log.Fatalf("Could not parse %s into uuid", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := UpdateUser(userId, request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

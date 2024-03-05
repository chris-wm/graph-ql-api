package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Type       string  `json:"type" validate:"omitempty,templateType"`
	Visibility string  `json:"visibility" validate:"omitempty,visibility"`
	Content    *string `json:"content"`
	Name       string  `json:"name"`
}

func Update(c *gin.Context) {

	id := c.Param("id")

	// Validate request
	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template, err := UpdateTemplate(id, request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, template)
}

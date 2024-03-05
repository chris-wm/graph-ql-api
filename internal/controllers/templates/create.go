package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Type       string  `json:"type" validate:"required,templateType"`
	Visibility string  `json:"visibility" validate:"required,visibility"`
	Content    *string `json:"content" validate:"required"`
	Name       string  `json:"name"`
}

func Create(c *gin.Context) {

	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	template, err := CreateTemplate(request)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, template)
}

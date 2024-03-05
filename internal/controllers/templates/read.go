package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {

	id := c.Param("id")

	template, err := ReadTemplate(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, template)
}

package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteResponse struct {
	ResourceId string `json:"resourceId"`
	Message    string `json:"message"`
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := DeleteTemplate(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := DeleteResponse{}
	response.ResourceId = id
	response.Message = "Template has been successfully deleted"

	c.JSON(http.StatusOK, response)
}

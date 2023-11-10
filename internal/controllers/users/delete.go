package users

import (
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteResponse struct {
	ResourceId string `json:"resourceId"`
	Message    string `json:"message"`
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	userId, err := uuid.Parse(id)

	if err != nil {
		log.Fatalf("Could not parse %s into uuid", id)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = DeleteUser(userId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := DeleteResponse{}
	response.ResourceId = id
	response.Message = "User has been successfully deleted"

	c.JSON(http.StatusOK, response)
}

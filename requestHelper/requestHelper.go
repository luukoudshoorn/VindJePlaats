package requestHelper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Response dispatcher from: https://levelup.gitconnected.com/sending-http-error-codes-with-golang-and-gin-gonic-d915d1dd0166
type Response struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

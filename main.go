package main

import (
	"github.com/gin-gonic/gin"
	"oudshoorn.dev/mijnwoonplaats/geospatial"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	geospatial.SetupRouter(router)
	return router
}

func main() {
	router := SetupRouter()
	router.Run("localhost:8080")
}

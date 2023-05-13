package geospatial

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type coordinate struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type boundingBox struct {
	TopLeft     coordinate `json:"topleft"`
	BottomRight coordinate `json:"bottomright"`
}

type polygon struct {
	Vertices []coordinate `json:"vertices"`
}

type area struct {
	Id       int64   `json:"id"`
	Boundary polygon `json:"boundary"`
}

var areas = []area{}

func SetupRouter(router *gin.Engine) {
	router.GET("/areas", GetAreas)
	router.POST("/areas", PostAreas)
}

func GetAreas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, areas)
}

func PostAreas(c *gin.Context) {
	var newArea area

	if err := c.BindJSON(&newArea); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Could not parse JSON object"})
		return
	}

	if len(newArea.Boundary.Vertices) < 3 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Area boundary should contain at least 3 vertices"})
		return
	}

	newArea.Id = int64(len(areas) + 1)

	areas = append(areas, newArea)
	c.IndentedJSON(http.StatusCreated, newArea)
}

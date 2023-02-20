package geospatial

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"oudshoorn.dev/mijnwoonplaats/requestTesting"
)

var testBoundingBoxes = []boundingBox{
	{TopLeft: coordinate{Lon: 6.0165, Lat: 52.0114}, BottomRight: coordinate{Lon: 6.0445, Lat: 51.9982}}, //'Rheden'
	{TopLeft: coordinate{Lon: 5.9507, Lat: 52.0113}, BottomRight: coordinate{Lon: 5.9882, Lat: 51.9977}}, //'Velp noord'
	{TopLeft: coordinate{Lon: 6.0536, Lat: 52.0244}, BottomRight: coordinate{Lon: 6.0687, Lat: 52.0167}}, //'De Steeg'
	{TopLeft: coordinate{Lon: 6.0793, Lat: 52.0371}, BottomRight: coordinate{Lon: 6.0957, Lat: 52.0277}}, //'Ellecom'
	{TopLeft: coordinate{Lon: 6.0832, Lat: 52.0700}, BottomRight: coordinate{Lon: 6.1203, Lat: 52.0390}}, //'Dieren'
	{TopLeft: coordinate{Lon: 5.9282, Lat: 51.9902}, BottomRight: coordinate{Lon: 5.9396, Lat: 51.9850}}, //'Arnhem'
	{TopLeft: coordinate{Lon: 6.1144, Lat: 52.4917}, BottomRight: coordinate{Lon: 6.1240, Lat: 52.4837}}, //'Zwolle'
	{TopLeft: coordinate{Lon: 4.8861, Lat: 52.1160}, BottomRight: coordinate{Lon: 4.9015, Lat: 52.1065}}, //'Kamerik'
	{TopLeft: coordinate{Lon: 6.8433, Lat: 52.2378}, BottomRight: coordinate{Lon: 6.8553, Lat: 52.2272}}, //'Enschede'
	{TopLeft: coordinate{Lon: 5.8418, Lat: 51.8876}, BottomRight: coordinate{Lon: 5.8532, Lat: 51.8813}}, //'Oosterhout'
}

var testArea = area{Id: -1, Boundary: polygon{Vertices: []coordinate{{Lon: 6.0165, Lat: 52.0114}, {Lon: 6.0445, Lat: 52.0114}, {Lon: 6.0445, Lat: 51.9982}, {Lon: 6.0165, Lat: 51.9982}}}}

func TestAreas(t *testing.T) {
	router := gin.Default()
	SetupRouter(router)

	//Test that areas starts empty and returns a valid response
	assert.Equal(t, requestTesting.GetTest(t, router, "/areas", 200), "[]")

	//Test an empty and an invalid area
	requestTesting.PostTest(t, router, "/areas", bytes.NewBuffer([]byte{}), 400)
	invalidArea, _ := json.Marshal(testBoundingBoxes[0])
	requestTesting.PostTest(t, router, "/areas", bytes.NewBuffer(invalidArea), 400)

	//Post an area
	areaPostJsonInput, _ := json.Marshal(testArea)
	response := requestTesting.PostTest(t, router, "/areas", bytes.NewBuffer(areaPostJsonInput), 201)
	testArea.Id = 1
	var areaResponse area
	json.Unmarshal([]byte(response), &areaResponse)
	assert.Equal(t, testArea, areaResponse)

	//Check that areas now contains 1 area
	response = requestTesting.GetTest(t, router, "/areas", 200)
	var areaResponseArray []area
	json.Unmarshal([]byte(response), &areaResponseArray)
	assert.Equal(t, []area{testArea}, areaResponseArray)
}

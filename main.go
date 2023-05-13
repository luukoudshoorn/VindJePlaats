package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"oudshoorn.dev/mijnwoonplaats/bagImporter"
	"oudshoorn.dev/mijnwoonplaats/geospatial"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	geospatial.SetupRouter(router)
	return router
}

func StartServer(port int) {
	router := SetupRouter()
	router.Run(fmt.Sprintf("localhost:%d", port))
}

func main() {
	importBag := flag.String("import-bag", "", "Import the bag extract using the filename provided as argument. File should be a downloaded zip BAG extract from the kadaster site.")
	runServer := flag.Bool("run", false, "Run the server")
	port := flag.Int("port", 8080, "The port that the server will listen on")

	flag.Parse()

	Run(*runServer, *port, *importBag)
}

func Run(runServer bool, port int, importBag string) {
	//Initialize the database
	//TODO

	//Run the desired imports
	if importBag != "" {
		bagImporter.ImportBag(importBag)
	}

	//Run the server
	if runServer {
		//Check if all data is now available
		//TODO

		//Finally start the server
		StartServer(port)
	}
}

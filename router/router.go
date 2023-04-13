package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Initialize the routes

	initializeRoutes(router)

	// Start serving the application
	router.Run(":8080") // listen and serve on 0.0.0.0:8080

}

package router

import (
	"firstproject/API/handler"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
)

// Unused
func createRateLimiter() *limiter.Limiter {
	// Set limit to 1 request per second per client
	requestLimit := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})

	// Set message to be displayed when client hits the limit
	requestLimit.SetMessage("Too many requests, please try again later.")

	// Return the rate limiter
	return requestLimit
}

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{

		// Show opening
		v1.GET("/opening", handler.CreateOpeningHandler)
		v1.POST("/opening", handler.ShowOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}

}

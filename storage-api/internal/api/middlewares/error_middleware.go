package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Using defer to recover from panics
		defer func() {
			r := recover()

			if r != nil {
				// Log the panic details
				log.Printf("Recovered from panic: %v", r)
				// Send a generic error message to the client
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong, please try again later",
				})
			}
		}()

		// Continue with the request handling
		ctx.Next()
	}
}

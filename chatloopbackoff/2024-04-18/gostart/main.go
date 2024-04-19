package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

const defaultMessage = "Hello! OpenFeature says NO!!"
const newWelcomeMessage = "Hello, welcome to this OpenFeature-enabled website!"

func main() {
	// Use flagd as the OpenFeature provider
	err := openfeature.SetProviderAndWait(flagd.NewProvider())
	if err != nil {
		// If an error occurs, log it and exit
		log.Fatalf("Failed to set the OpenFeature provider: %v", err)
	}

	// Initialize OpenFeature client
	client := openfeature.NewClient("GoStartApp")

	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("/hello", func(c *gin.Context) {

		// Evaluate welcome-message feature flag
		welcomeMessage, _ := client.BooleanValue(
			context.Background(), "welcome-message", false, openfeature.EvaluationContext{},
		)

		if welcomeMessage {
			c.JSON(http.StatusOK, newWelcomeMessage)
			return
		} else {
			c.JSON(http.StatusOK, defaultMessage)
			return
		}
	})

	engine.Run()
}

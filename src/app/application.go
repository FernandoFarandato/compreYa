package app

import (
	"compreYa/src/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

const port = ":8080"

func Start() {
	router := createCustomRouter()

	handlers := dependencies.Start()

	configureURLMapping(router, handlers)

	router.Run(port)
}

func createCustomRouter() *gin.Engine {
	router := gin.Default()
	// ping handler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// add middleware
	// router.Use(midleware)

	return router
}

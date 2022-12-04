package app

import (
	"compreYa/src/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func configureURLMapping(router *gin.Engine, handlers *dependencies.HandlerContainer) {
	authGroup := router.Group("compreYa/auth")
	authGroup.POST("/signup", handlers.SignUp.Handle)
	authGroup.POST("/login", handlers.Login.Handle)
	authGroup.GET("/validate", handlers.AuthValidation.Handle)

	storeGroup := router.Group("compreYa/store")
	storeGroup.GET("/create", handlers.AuthValidation.Handle, handlers.CreateStore.Handle)
	storeGroup.GET("/validate/name", handlers.SignUp.Handle)
}

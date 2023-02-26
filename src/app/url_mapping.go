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
	authGroup.POST("/recover_password/request", handlers.RecoverPasswordRequest.Handle)
	authGroup.POST("/recover_password/change", handlers.RecoverPasswordChange.Handle)

	userInfoGroup := router.Group("compreYa/user")
	userInfoGroup.POST("/change/password", handlers.Login.Handle) // pending

	storeGroup := router.Group("compreYa/store", handlers.AuthValidation.Handle)
	storeGroup.POST("/create", handlers.CreateStore.Handle)
	storeGroup.POST("/delete/:url_name", handlers.DeleteStore.Handle)
	storeGroup.GET("/validate/name", handlers.SignUp.Handle)
}

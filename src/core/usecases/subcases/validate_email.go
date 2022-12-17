package subcases

import "github.com/gin-gonic/gin"

type ValidateEmail interface {
	Execute(c *gin.Context, email string)
}

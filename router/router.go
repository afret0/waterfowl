package router

import "strings"

func RouterTem(svr string) string {
	t := `
package router

import (
	"github.com/gin-gonic/gin"
	"sample/infrastructure/middleware"
	"sample/infrastructure/router"
	"sample/handler"
)

func RegisterRouter(e *gin.Engine) {
	r := router.GetRouter(e)

	svr := handler.GetService()
	//r.Use(middleware.AuthMiddleware())
	r.POST("/ping", svr.Ping)

	sampleRouter := r.Group("/sample")
	sampleRouter.Use(middleware.AuthMiddleware())
	sampleRouter.POST("/ping", svr.Ping)

	internal := r.Group("/internal")
	internal.POST("/ping", svr.Ping)

	r.RegisterRouter()
}

	
`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}

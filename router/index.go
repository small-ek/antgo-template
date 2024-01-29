package router

import (
	"antgo-template/app/http/index"
	"github.com/gin-gonic/gin"
)

func IndexRoute(Router *gin.RouterGroup) {
	Router.GET("/", new(index.IndexController).Index)
}

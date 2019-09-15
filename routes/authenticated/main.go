package authenticatedroutes

import (
	"github.com/gin-gonic/gin"
)

// Bootstrap the routes
func Bootstrap(router *gin.RouterGroup) {
	router.GET("/user", userPage)
	router.GET("/", homepage)
}

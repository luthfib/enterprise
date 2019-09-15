package publicroutes

import "github.com/gin-gonic/gin"

// Bootstrap the routes
func Bootstrap(router *gin.RouterGroup) {
	//router.GET("/", homepage)
	router.GET("/login", login)
	router.GET("/callback", callback)
	router.GET("/logout", logout)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

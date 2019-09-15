package authenticatedroutes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func homepage(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("UserID")
	fmt.Println("Homepage userID", userID)

	// c.HTML(http.StatusOK, templateName, data)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home Page",
		//"payload": articles},
	},
	)
}

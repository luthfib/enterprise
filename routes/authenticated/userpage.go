package authenticatedroutes

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func userPage(c *gin.Context) {
	session := sessions.Default(c)
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("users")
	fmt.Println(collection)
	fmt.Println("user", session.Get("user"))
	// c.HTML(http.StatusOK, templateName, data)
	c.HTML(http.StatusOK, "user.html", gin.H{
		"title": "Home Page",
		//"payload": articles},
	},
	)
}

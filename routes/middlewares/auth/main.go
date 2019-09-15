package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Required() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("this entered")
		session := sessions.Default(c)
		fmt.Println("sessio", session)
		profile := session.Get("profile")

		if profile == nil {
			fmt.Println("Entered the profiel statement")
			c.Redirect(http.StatusSeeOther, "/login")
		} else {
			c.Next()
		}
	}
}

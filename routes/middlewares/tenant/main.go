package tenant

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("begin", t)
		session := sessions.Default(c)
		if session.Get("TenantID") != nil && session.Get("UserID") != nil {
			c.Next()
		}

		db := c.MustGet("DB").(*mongo.Client)
		users := db.Database("db1").Collection("users")

		if session.Get("profile") == nil {
			c.Redirect(http.StatusSeeOther, "/")
			
			return
		}
		profile := session.Get("profile").(map[string]interface{})
		email := profile["name"]

		result := bson.Raw{}
		filter := bson.M{"Email": email.(string)}
		err := users.FindOne(c, filter).Decode(&result)
		if err != nil {
			c.JSON(http.StatusCreated, gin.H{"Error": "Unable to find user"})
			return
		}

		var tenantID string
		var userID string

		result.Validate()
		userIDRaw := result.Lookup("UserID")
		userIDRaw.Unmarshal(&userID)
		tenantIDRaw := result.Lookup("TenantID")
		tenantIDRaw.Unmarshal(&tenantID)
		session.Set("UserID", userID)
		session.Set("TenantID", tenantID)
		err = session.Save()
		if err != nil {
			c.JSON(http.StatusCreated, gin.H{"Error": "Unable to set session values"})
			return
		}
		fmt.Println("end", t)
		c.Next()

	}
}

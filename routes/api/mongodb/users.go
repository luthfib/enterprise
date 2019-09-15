package mongodb

import (
	"context"
	"enterprise/models"
	utils "enterprise/utilities"
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	validator "gopkg.in/go-playground/validator.v9"
	// "gopkg.in/go-playground/validator.v9"
)

func CreateUser(c *gin.Context) {
	session := sessions.Default(c)
	tenantID := session.Get("TenantID")
	user := models.User{
		TenantID:  tenantID.(string),
		UserID:    utils.RandStr(20),
		FirstName: c.Query("FirstName"),
		LastName:  c.Query("LastName"),
		Email:     c.Query("Email"),
		TimeZone:  c.Query("TimeZone"),
		Country:   c.Query("Country"),
		AvatarURL: c.Query("AvatarURL"),
	}

	validator := c.MustGet("Validator").(*validator.Validate)
	err := validator.Struct(user)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"Error": err.Error()})
		return
	}

	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("users")
	res, err := collection.InsertOne(c, structs.Map(user))

	fmt.Println(res)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusCreated, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": res})
}

func DeleteUser(c *gin.Context) {
	session := sessions.Default(c)
	tenantID := session.Get("TenantID").(string)
	userID := c.Query("UserID")

	if utils.ContainsEmpty(tenantID, userID) {
		c.JSON(http.StatusCreated, gin.H{"Error": "TenantID or UserID is Empty"})
		return
	}
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("users")

	filter := bson.D{
		{"TenantID", tenantID},
		{"UserID", userID},
	}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"User Deleted" + tenantID: res})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"article": "not ok"})
}

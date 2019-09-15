package mongodb

import (
	"enterprise/models"
	utils "enterprise/utilities"
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
	validator "gopkg.in/go-playground/validator.v9"
)

func CreateTenant(c *gin.Context) {
	tenant := models.Tenant{
		TenantID:    utils.RandStr(20),
		FirstName:   c.Query("FirstName"),
		LastName:    c.Query("LastName"),
		LicenseType: c.Query("LicenseType"),
	}
	validator := c.MustGet("Validator").(*validator.Validate)
	err := validator.Struct(tenant)
	if err != nil {
		fmt.Println("OK SO ERROR IS HERE", err)
		c.JSON(http.StatusCreated, gin.H{"Error": err.Error()})
		return
	}
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("tenants")
	res, err := collection.InsertOne(c, structs.Map(tenant))

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusCreated, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": res})
}

package mongodb

import (
	"enterprise/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/mongodb/mongo-go-driver/bson"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/fatih/structs"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateProject(c *gin.Context) {
	session := sessions.Default(c)
	tenantID := session.Get("TenantID")
	userID := session.Get("UserID")
	project := models.Project{
		TenantID:      tenantID.(string),
		ProjectID:     c.Query("ProjectID"),
		Name:          c.Query("Name"),
		Description:   c.Query("Description"),
		CreatorUserID: userID.(string),
	}

	validator := c.MustGet("Validator").(*validator.Validate)
	err := validator.Struct(project)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"Error": err.Error()})
		return
	}
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("projects")
	res, err := collection.InsertOne(c, structs.Map(project))

	res, err = collection.InsertOne(c, bson.M{
		"TenantID":  project.TenantID,
		"ProjectID": project.ProjectID,
		"Name":      project.Name,
	})

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusCreated, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": res})
}

func DeleteProject(c *gin.Context) {

	docToDelete := bson.M{"TenantID": c.Query("TenantID"),
		"Name": c.Query("Name")}

	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("projects")
	res, err := collection.DeleteOne(c, docToDelete)

	fmt.Println(res)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusCreated, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": "a"})
}

func UpdateProject(c *gin.Context) {
	projectID := c.Query("ProjectID")
	value := c.Query("value")
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("projects")
	fmt.Print(collection)

	if len(value) == 0 {
		c.JSON(http.StatusCreated, gin.H{"Error": "Value Is Required"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"projectID": projectID, "value": value})
}

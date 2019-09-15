package mongodb

import (
	"context"
	"enterprise/models"
	utils "enterprise/utilities"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// objectid.ObjectID
func CreateTask(c *gin.Context) {
	session := sessions.Default(c)
	tenantID := session.Get("TenantID")
	userID := session.Get("UserID")
	task := models.Task{
		TenantID:      tenantID.(string),
		TaskID:        primitive.NewObjectID(),
		TaskType:      c.Query("TaskType"),
		CreatorUserID: userID.(string),
		Name:          c.Query("TaskType"),
		Description:   c.Query("TaskType"),
		CreatedTime:   time.Now(),
		Assignee:      c.Query("TaskType"),
		TimeEstimate:  c.Query("TaskType"),
		IsBlockedBy:   c.Query("TaskType"),
		Blocks:        c.Query("TaskType"),
		StartDate:     c.Query("TaskType"),
		EndDate:       c.Query("TaskType"),
		Status:        c.Query("TaskType"),
		Priority:      c.Query("TaskType"),
		IsRejected:    c.Query("TaskType"),
		Rejector:      c.Query("TaskType"),
		ParentTaskID:  c.Query("TaskType"),
		Label:         c.Query("TaskType"),
	}

	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("tasks")
	res, err := collection.InsertOne(c, structs.Map(task))

	fmt.Println(res)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusCreated, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Success": res})
}

func DeleteTask(c *gin.Context) {
	session := sessions.Default(c)
	tenantID := session.Get("TenantID").(string)
	userID := session.Get("UserID").(string)
	taskID := c.Query("TaskID")

	if utils.ContainsEmpty(tenantID, userID, taskID) {
		c.JSON(http.StatusCreated, gin.H{"Error": "A value is empty"})
		return
	}
	db := c.MustGet("DB").(*mongo.Client)
	collection := db.Database("db1").Collection("tasks")

	filter := bson.D{
		{"TenantID", tenantID},
		{"CreatorUserID", userID},
		{"TaskID", taskID},
	}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Task Deleted": res})
}

func UpdateTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"article": "not ok"})
}

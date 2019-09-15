package api

import (
	"enterprise/routes/api/mongodb"

	"github.com/gin-gonic/gin"
)

// Bootstrap the routes
func Bootstrap(router *gin.RouterGroup) {
	router.GET("/CreateTenant", mongodb.CreateTenant)

	router.GET("/CreateUser", mongodb.CreateUser)
	router.GET("/DeleteUser", mongodb.DeleteUser)
	router.GET("/UpdateUser", mongodb.UpdateUser)

	router.GET("/CreateProject", mongodb.CreateProject)
	router.GET("/DeleteProject", mongodb.DeleteProject)
	router.GET("/UpdateProject", mongodb.UpdateProject)

	router.GET("/CreateTask", mongodb.CreateTask)
	router.GET("/DeleteTask", mongodb.DeleteTask)
	router.GET("/UpdateTask", mongodb.UpdateTask)

	router.GET("/AddCommentToTask", mongodb.AddCommentToTask)
	router.GET("/DeleteCommentOnTask", mongodb.DeleteCommentOnTask)
	router.GET("/UpdateCommentOnTask", mongodb.UpdateCommentOnTask)

}

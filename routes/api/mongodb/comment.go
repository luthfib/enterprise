package mongodb

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCommentToTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"article": "not ok"})
}

func DeleteCommentOnTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"article": "Deleted"})
}

func UpdateCommentOnTask(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"article": "not ok"})
}

package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetComments(c *gin.Context) {

	var comments []models.Comment

	result := db.DB.Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func GetCommentByID(c *gin.Context) {

	var comment models.Comment

	id := c.Param("id")

	result := db.DB.First(&comment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

func CreateComment(c *gin.Context) {

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	type CreateCommentInput struct {
		Content string `form:"content" binding:"required"`
		PostID  uint   `form:"post_id" binding:"required"`
	}

	var input CreateCommentInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	comment := models.Comment{Content: input.Content, UserID: userId, PostID: input.PostID}

	result := db.DB.Create(&comment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, comment)

}

func UpdateComment(c *gin.Context) {

	type UpdateCommentInput struct {
		Content string `form:"content" binding:"required"`
	}

	var input UpdateCommentInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id := c.Param("id")

	var comment models.Comment

	result := db.DB.First(&comment, id)

	if result.Error != nil {

		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Content != "" {
		comment.Content = input.Content
	}

	result = db.DB.Save(&comment)

	if result.Error != nil {

		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

func DeleteComment(c *gin.Context) {

	id := c.Param("id")

	var comment models.Comment

	result := db.DB.First(&comment, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&comment)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, comment)

}

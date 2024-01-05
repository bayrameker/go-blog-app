package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {

	var posts []models.Post

	result := db.DB.Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func GetPostByID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func CreatePost(c *gin.Context) {

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	type CreatePostInput struct {
		Title       string `form:"title" binding:"required"`
		Description string `form:"description" binding:"required"`
		Content     string `form:"content" binding:"required"`
	}

	var input CreatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	post := models.Post{
		Title:       input.Title,
		Description: input.Description,
		Content:     input.Content,
		AuthorID:    userId,
	}

	result := db.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func UpdatePost(c *gin.Context) {

	type UpdatePostInput struct {
		Title       string `form:"title"`
		Description string `form:"description"`
		Content     string `form:"content"`
	}

	var input UpdatePostInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var post models.Post

	id := c.Param("id")

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Title != "" {
		post.Title = input.Title
	}

	if input.Description != "" {
		post.Description = input.Description
	}

	if input.Content != "" {
		post.Content = input.Content
	}

	result = db.DB.Save(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&post)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, post)

}

func GetCommentsByPostID(c *gin.Context) {

	id := c.Param("id")

	var comments []models.Comment

	result := db.DB.Where("post_id = ?", id).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})

}

func GetAuthorByPostID(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var author models.User

	result = db.DB.First(&author, post.AuthorID)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, author)
}

package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"example.com/blog-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func GetMyCredentials(c *gin.Context) {

	userId := c.GetUint("userId")

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	userId := c.GetUint("userId")

	type UpdateMyCredentialsInput struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateMyCredentialsInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Email != "" {
		user.Email = input.Email
	}

	if input.Password != "" {
		hashedPassword := utils.HashPassword(input.Password)
		user.Password = hashedPassword
	}

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyPosts(c *gin.Context) {

	userId := c.GetUint("userId")

	var posts []models.Post

	result := db.DB.Where("author_id = ?", userId).Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

func GetMyComments(c *gin.Context) {

	userId := c.GetUint("userId")

	var comments []models.Comment

	result := db.DB.Where("user_id = ?", userId).Find(&comments)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})

}

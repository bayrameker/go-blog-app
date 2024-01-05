package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"example.com/blog-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	type SignIn struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignIn

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	isPasswordTrue := utils.ComparePasswords(input.Password, user.Password)

	if !isPasswordTrue {
		c.Status(http.StatusUnauthorized)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func SignUp(c *gin.Context) {

	type SignUpInput struct {
		Username string `form:"username" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignUpInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var existingUser models.User

	result := db.DB.Where("email = ? OR username = ?", input.Email, input.Username).First(&existingUser)

	if result.Error == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hash := utils.HashPassword(input.Password)

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hash,
	}

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

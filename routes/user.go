package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"example.com/blog-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []models.User

	result := db.DB.Find(&users)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})

}

func GetUserByID(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func CreateUser(c *gin.Context) {

	type CreateUserInput struct {
		Username string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
		RoleID   uint   `form:"role_id" binding:"required"`
	}

	var input CreateUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hashedPassword := utils.HashPassword(input.Password)

	user := models.User{
		Username: input.Username,
		Email:    input.Username,
		Password: hashedPassword,
	}

	result := db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, user)

}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	type UpdateUserInput struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateUserInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.First(&user, id)

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

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	result := db.DB.First(&user, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {

	var roles []models.Role

	result := db.DB.Find(&roles)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": roles,
	})

}

func GetRoleByID(c *gin.Context) {

	id := c.Param("id")

	var role models.Role

	result := db.DB.First(&role, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, role)

}

func CreateRole(c *gin.Context) {

	type CreateRoleInput struct {
		Name string `form:"name" binding:"required"`
	}

	var input CreateRoleInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	role := models.Role{Name: input.Name}

	result := db.DB.Create(&role)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, role)

}

func UpdateRole(c *gin.Context) {

	type UpdateRoleInput struct {
		Name string `form:"name" binding:"required"`
	}

	var input UpdateRoleInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var role models.Role

	id := c.Param("id")

	result := db.DB.First(&role, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		role.Name = input.Name
	}

	result = db.DB.Save(&role)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, role)

}

func DeleteRole(c *gin.Context) {

	id := c.Param("id")

	var role models.Role

	result := db.DB.First(&role, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&role)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, role)

}

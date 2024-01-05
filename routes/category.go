package routes

import (
	"net/http"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {

	var categories []models.Category

	result := db.DB.Find(&categories)

	if result.Error != nil {
		c.Status(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
	})

}

func GetCategoryByID(c *gin.Context) {

	var category models.Category

	id := c.Param("id")

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func CreateCategory(c *gin.Context) {

	type CreateCategoryInput struct {
		Name string `form:"name" binding:"required"`
	}

	var input CreateCategoryInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	category := models.Category{Name: input.Name}

	result := db.DB.Create(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func UpdateCategory(c *gin.Context) {

	type UpdateCategoryInput struct {
		Name string `form:"name" binding:"required"`
	}

	var input UpdateCategoryInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var category models.Category

	id := c.Param("id")

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Name != "" {
		category.Name = input.Name
	}

	result = db.DB.Save(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func DeleteCategory(c *gin.Context) {

	var category models.Category

	id := c.Param("id")

	result := db.DB.First(&category, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&category)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, category)

}

func GetPostsByCategoryID(c *gin.Context) {

	id := c.Param("id")

	var posts []models.Post

	result := db.DB.Where("category_id = ?", id).Find(&posts)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})

}

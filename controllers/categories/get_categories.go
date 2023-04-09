package categories

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

// @Summary set Get Categories
// @Tags    categories
// @Router  /categories [get]
// @Accept 	json
// @Success 200           {array} 	models.Category
// @Failure 400
func GetCategories(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Category
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

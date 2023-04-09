package foods

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

// @Summary set Get foods
// @Tags    foods
// @Router  /foods [get]
// @Accept 	json
// @Success 200           {array} 	models.Food
// @Param        categoryName   query      string  false  "Category Name"
// @Param        subCategoryName   query      string  false  "Sub Category Name"
// @Failure 400
func GetFoods(c *gin.Context) {
	var err error
	categoryName := c.Query("categoryName")
	subCategoryName := c.Query("categoryName")
	db := database.GetDatabase()
	var foods []models.Food
	if categoryName != "" {
		if subCategoryName != "" {
			err = db.Where("category_name = ?", categoryName).Where("sub_category_name = ?", subCategoryName).Find(&foods).Error
		} else {
			err = db.Where("category_name = ?", categoryName).Find(&foods).Error
		}
	} else {
		err = db.Find(&foods).Error
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find products: " + err.Error(),
		})
		return
	}

	c.JSON(200, foods)
}

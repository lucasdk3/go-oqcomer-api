package foods

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

// @Summary set Sort Food
// @Tags    foods
// @Router  /foods/sort/{id} [get]
// @Accept 	json
// @Param        id   path      int  true  "Food id"
// @Param        categoryName   query      string  false  "Category Name"
// @Param        subCategoryName   query      string  false  "Sub Category Name"
// @Success 200           {object} models.Food
// @Failure 400
func SortFood(c *gin.Context) {
	var err error
	categoryName := c.Query("categoryName")
	subCategoryName := c.Query("subCategoryName")
	db := database.GetDatabase()
	var food models.Food
	if categoryName != "" {
		if subCategoryName != "" {
			err = db.Where("category_name = ?", categoryName).Where("sub_category_name = ?", subCategoryName).Order("RAND()").First(&food).Error
		} else {
			err = db.Where("category_name = ?", categoryName).Order("RAND()").First(&food).Error
		}
	} else {
		err = db.Order("RAND()").First(&food).Error
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, food)
}

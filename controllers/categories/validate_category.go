package categories

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

// @Summary set Get Category
// @Tags    categories
// @Router  /categories/{id} [get]
// @Accept 	json
// @Param        id   path      int  true  "Category id"
// @Success 200           {object} models.Category
// @Failure 400
func ValidateCategory(c *gin.Context, name string) error {

	db := database.GetDatabase()
	var category models.Category
	err := db.Where("name = ?", name).First(&category).Error

	if err != nil {
		return err
	}

	return nil
}

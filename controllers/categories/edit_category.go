package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

type EditCategoryRequestBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// @Summary set Edit Category
// @Tags    categories
// @Router  /categories [put]
// @Accept 	json
// @Param   payload   body    EditCategoryRequestBody    false  "Category name"
// @Success 200           {object} models.Category
// @Failure 400
func EditCategory(ctx *gin.Context) {
	body := EditCategoryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category models.Category

	db := database.GetDatabase()

	err := db.First(&category, body.ID).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot edit category: " + err.Error(),
		})
		return
	}

	category.Name = body.Name
	category.ID = body.ID

	err = db.Save(&category).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, category)
}

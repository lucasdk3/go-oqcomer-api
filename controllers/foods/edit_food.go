package foods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

type EditFoodRequestBody struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	CategoryName    string `json:"categoryName"`
	SubCategoryName string `json:"subCategoryName"`
}

// @Summary set Edit Food
// @Tags    foods
// @Router  /foods [put]
// @Accept 	json
// @Param   payload   body    EditFoodRequestBody    false  "Food name"
// @Success 200           {object} models.Food
// @Failure 400
func EditFood(ctx *gin.Context) {
	body := EditFoodRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var food models.Food

	db := database.GetDatabase()

	err := db.First(&food, body.ID).Error

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot edit category: " + err.Error(),
		})
		return
	}

	food.ID = body.ID
	food.Name = body.Name
	food.CategoryName = body.CategoryName
	food.SubCategoryName = body.SubCategoryName

	err = db.Save(&food).Error
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	ctx.JSON(200, food)
}

package foods

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/controllers/categories"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

type AddFoodRequestBody struct {
	Name            string `json:"name"`
	CategoryName    string `json:"categoryName"`
	SubCategoryName string `json:"subCategoryName"`
}

// @Summary set Add Food
// @Tags    foods
// @Router  /foods [post]
// @Accept 	json
// @Param   payload   body    AddFoodRequestBody    false  "Food"
// @Success 201           {object} models.Food
// @Failure 400
func AddFood(ctx *gin.Context) {
	body := AddFoodRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var Food models.Food

	Food.Name = body.Name

	if err := categories.ValidateCategory(ctx, body.CategoryName); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else {
		Food.CategoryName = body.CategoryName
	}

	if err := categories.ValidateCategory(ctx, body.SubCategoryName); err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else {
		Food.SubCategoryName = body.SubCategoryName
	}

	db := database.GetDatabase()

	if result := db.Create(&Food); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &Food)
}

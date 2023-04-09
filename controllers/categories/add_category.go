package categories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

type AddCategoryRequestBody struct {
	Name string `json:"name"`
}

// @Summary set Add Category
// @Tags    categories
// @Router  /categories [post]
// @Accept 	json
// @Param   payload   body    AddCategoryRequestBody    false  "Category name"
// @Success 201           {object} models.Category
// @Failure 400
func AddCategory(ctx *gin.Context) {
	body := AddCategoryRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var category models.Category

	category.Name = body.Name

	db := database.GetDatabase()

	if result := db.Create(&category); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &category)
}

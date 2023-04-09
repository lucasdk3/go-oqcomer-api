package foods

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/database"
	"github.com/lucasdk3/go-oqcomer-api/models"
)

// @Summary set Delete Food
// @Tags    foods
// @Router  /foods [delete]
// @Accept 	json
// @Param   id    query   uint  false  "id"
// @Success 204           {object} models.Food
// @Failure 400
func DeleteFood(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Food{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

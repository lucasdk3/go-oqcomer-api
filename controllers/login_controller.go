package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/maui-oqcomer-api/database"
	"github.com/lucasdk3/maui-oqcomer-api/models"
	"github.com/lucasdk3/maui-oqcomer-api/services"
)

// @Summary set login
// @Tags    login
// @Router  /login [post]
// @Accept 	json
// @Param   payload   body    models.Login    false  "Username e Pin"
// @Success 200           {object} models.Login
// @Failure 400
func Login(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Login
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("username = ?", p.Username).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})
		return
	}

	if user.Pin != services.SHA256Encoder(p.Pin) {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.JWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})

}

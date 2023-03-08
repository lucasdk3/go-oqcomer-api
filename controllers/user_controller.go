package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/maui-oqcomer-api/database"
	"github.com/lucasdk3/maui-oqcomer-api/models"
	"github.com/lucasdk3/maui-oqcomer-api/services"
)

// @Summary set user
// @Tags    user
// @Router  /user [post]
// @Accept 	json
// @Param   payload   body    models.User    true  "Dados do usuário"
// @Success 204           {object} models.User
// @Failure 400
func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	p.Pin = services.SHA256Encoder(p.Pin)

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary get user
// @Tags    user
// @Router  /user [get]
// @Accept 	json
// @Success 200           {object} models.User
// @Failure 400
// @Security BearerAuth
func GetUser(c *gin.Context) {
	const Bearer_schema = "Bearer "
	header := c.GetHeader("Authorization")

	tokenString := header[len(Bearer_schema):]
	userId, err := services.JWTService().GetJWTClaim(tokenString)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "User is invalid",
		})
		return
	}

	db := database.GetDatabase()
	var p models.User
	err = db.First(&p, userId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user by id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

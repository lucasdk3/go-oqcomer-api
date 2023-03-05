package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/maui-oqcomer-api/database"
	"github.com/lucasdk3/maui-oqcomer-api/models"
	"github.com/lucasdk3/maui-oqcomer-api/services"
)

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
			"error": "cannot create book: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

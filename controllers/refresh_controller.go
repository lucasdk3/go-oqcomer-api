package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lucasdk3/go-oqcomer-api/services"
)

// @Summary set refresh
// @Tags    refresh
// @Router  /refresh [post]
// @Accept 	json
// @Param   refresh    query   string  false  "refresh"
// @Success 200           {object} models.Login
// @Failure 400
func Refresh(c *gin.Context) {
	refreshToken := c.Query("refresh")

	token, err := services.JWTService().ValidateRefreshToken(refreshToken)

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	if token.Valid {
		userId, claimErr := services.JWTService().GetJWTRefreshClaim(refreshToken)

		if claimErr != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		token, err := services.JWTService().GenerateToken(userId)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"access":  token.AccessToken,
			"refresh": token.RefreshToken,
		})
	} else {
		c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"user-service/service"
)

func tokenValidation(c *gin.Context) {
	token := c.Query("token")
	tokenService := service.TokenService{}
	clam, err := tokenService.Decode(token)
	if err != nil {
		log.Infof("/token validation can't decode token Error :  %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't decode token",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "valid token",
		"data":    clam.User,
	})
}

package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"user-service/db"
	"user-service/model"
	"user-service/service"
)

func createUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Warnf("/create user can't parse request body Error :  %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't parse request body",
			"data":    err.Error(),
		})

		return
	}

	if !user.IsValid() {
		log.Warnln("/create user invalid request")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})

		return
	}

	if err := db.Client.Create(&user); err != nil {
		log.Warnf("/create user can't save in database Error :  %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't save in database",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
	})
}

func login(c *gin.Context) {
	var reqUser model.User
	if err := c.BindJSON(&reqUser); err != nil {
		log.Warnf("/login can't parse request body Error :  %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't parse request body",
			"data":    err.Error(),
		})

		return
	}

	user, err := db.Client.GetByEmail(reqUser.Email)
	if err != nil || user.Password != reqUser.Password {
		log.Warnf("/login can't retrieve user form database Error :  %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Invalid user credentials",
			"data":    err.Error(),
		})

		return
	}

	tokenService := service.TokenService{}
	tokenStr, err := tokenService.Encode(user)
	if err != nil {
		log.Warnf("/login can't generate token Error :  %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't generate token",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful",
		"data":    tokenStr,
	})
}

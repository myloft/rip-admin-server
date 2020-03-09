package service

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/repository"
	"net/http"
	"strconv"
)

func Login(c *gin.Context) {
	password := controller.GetPassword(c.PostForm("username"))
	if password == c.PostForm("password") && password != "" {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": EncodeToken(c.PostForm("username"))})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
	}
}

func GetAllUsers(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": controller.GetAllUsers()})
}

func UpdateUser(c *gin.Context) {
	var user repository.User
	user.Name = c.Query("username")
	AdminBool, _ := strconv.ParseBool(c.Query("admin"))
	user.Admin = &AdminBool
	if controller.UpdateUser(user) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
	} else {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
	}
}
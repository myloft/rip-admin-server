package service

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/repository"
	"net/http"
)

func GetAllMenus(c *gin.Context) {
	var menu repository.Menu
	if VerifyToken(c) {
		menu.Id = -1
		GetMenus(&menu)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": menu.Children})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

func GetMenus(menu *repository.Menu) {
	if menu.Id == -1 {
		controller.GetChildrenMenus(menu)
	}
	for i, _ := range menu.Children {
		controller.GetChildrenMenus(&menu.Children[i])
		GetMenus(&menu.Children[i])
	}
}
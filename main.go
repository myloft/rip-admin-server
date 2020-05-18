package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/service"
)

func main() {
	// CORS 配置
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization"}

	router := gin.Default()
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/animes", service.GetAllAnimes)
		v1.GET("/animes/:anid", service.GetSingleAnime)
		v1.POST("/animes", service.AddNewAnime)
		v1.PUT("/animes", service.UpdateAnime)
		v1.DELETE("/animes/:anid", service.DeleteAnime)
		v1.GET("/animeinfo", service.GetAnimeInfo)
		v1.POST("/login", service.Login)
		v1.GET("/users", service.GetAllUsers)
		v1.PUT("/users", service.UpdateUser)
		v1.GET("/users/animes", service.GetUserAnimes)
		v1.PUT("/users/password", service.UpdatePassword)
		v1.GET("/menus", service.GetAllMenus)
	}
	router.Run(":8081")
}

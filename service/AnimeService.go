package service

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/gin-gonic/gin"
	"main/controller"
	"main/repository"
	"net/http"
	"os"
	"strconv"
)

func GetAllAnimes(c *gin.Context) {
	var animes []repository.Anime
	if VerifyToken(c) {
		// 查询数量与页数 无参默认全部
		query := c.Query("query")
		pagenum, _ := strconv.Atoi(c.Query("pagenum"))
		pagesize, _ := strconv.Atoi(c.Query("pagesize"))
		status := c.Query("status")
		if controller.GetAllAnimes(&animes, query, pagenum, pagesize, status) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": animes, "total": controller.GetAllAnimesNum()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

func GetSingleAnime(c *gin.Context) {
	var anime repository.Anime
	if VerifyToken(c) {
		if controller.GetSingleAnime(c.Param("anid"), &anime) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": anime})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusNotFound})
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

func GetUserAnimes(c *gin.Context) {
	var animes []repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		// 查询数量与页数 无参默认全部
		query := c.Query("query")
		pagenum, _ := strconv.Atoi(c.Query("pagenum"))
		pagesize, _ := strconv.Atoi(c.Query("pagesize"))
		status, _ := strconv.Atoi(c.Query("status"))
		if controller.GetUserAnimes(&animes, username, query, pagenum, pagesize, status) {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": animes, "total": controller.GetUserAnimesNum(username)})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}

func AddNewAnime(c *gin.Context) {
	var anime repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		anime.Anid, _ = strconv.Atoi(c.PostForm("Anid"))
		anime.Owner = username
		anime.Official_name = c.PostForm("Official_name")
		anime.Zh_name = c.PostForm("Zh_name")
		anime.Status, _ = strconv.Atoi(c.PostForm("Status"))
		controller.AddNewAnime(&anime)
		c.JSON(http.StatusAccepted, gin.H{"status": http.StatusAccepted})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

func DeleteAnime(c *gin.Context) {
	var anime repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		anime.Anid, _ = strconv.Atoi(c.Param("anid"))
		anime.Owner = username
		controller.DeleteAnime(&anime)
		c.JSON(http.StatusAccepted, gin.H{"status": http.StatusAccepted})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

type Aidb struct {
	Official_name string
	Zh_name       string
}

func GetAnimeInfo(c *gin.Context) {
	var aidb Aidb
	f, err := os.Open("anime-titles.xml")
	doc, err := xmlquery.Parse(f)
	query := fmt.Sprintf("//anime[@aid='%s']", c.Query("aid"))
	titles := xmlquery.FindOne(doc, query)
	if title := titles.SelectElement("//title[@xml:lang='zh-Hans'][@type='official']"); title != nil {
		aidb.Zh_name = title.InnerText()
	}
	if title := titles.SelectElement("//title[@xml:lang='ja']"); title != nil {
		aidb.Official_name = title.InnerText()
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": aidb})
	}
}
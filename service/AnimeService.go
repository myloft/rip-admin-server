package service

import (
	"encoding/json"
	"fmt"
	"main/controller"
	"main/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/antchfx/xmlquery"
	"github.com/gin-gonic/gin"
)

func GetAllAnimes(c *gin.Context) {
	var animes []repository.Anime
	if VerifyToken(c) {
		// 查询数量与页数 无参默认全部
		query := c.Query("query")
		pagenum, _ := strconv.Atoi(c.Query("pagenum"))
		pagesize, _ := strconv.Atoi(c.Query("pagesize"))
		status, _ := strconv.Atoi(c.Query("status"))
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
		anime.Anid, _ = strconv.Atoi(c.Param("anid"))
		if controller.GetSingleAnime(&anime) {
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
		if controller.IsAdmin(username) {
			GetAllAnimes(c)
		} else {
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
}

func AddNewAnime(c *gin.Context) {
	var anime repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		anime.Anid, _ = strconv.Atoi(c.PostForm("Anid"))
		anime.Owner = username
		anime.Official_name = c.PostForm("Official_name")
		anime.Zh_name = c.PostForm("Zh_name")
		anime.Status, _ = strconv.Atoi(c.DefaultPostForm("Status", "1"))
		controller.AddNewAnime(&anime)
		c.JSON(http.StatusAccepted, gin.H{"status": http.StatusAccepted})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

func UpdateAnime(c *gin.Context) {
	var anime repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		anime.Anid, _ = strconv.Atoi(c.PostForm("Anid"))
		anime.Owner = username
		anime.Official_name = c.PostForm("Official_name")
		anime.Zh_name = c.PostForm("Zh_name")
		anime.Status, _ = strconv.Atoi(c.DefaultPostForm("Status", "1"))
		controller.UpdateAnime(&anime)
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

func GetPublishedInfo(anime *repository.Anime) bool {
	resp, err := http.Get("https://snow.hacg.top/api.php?anid=" + strconv.Itoa(anime.Anid))
	json.NewDecoder(resp.Body).Decode(anime)
	anime.Status = 2
	if err != nil {
		return false
	} else {
		return true
	}
}

func SetPublished(c *gin.Context) {
	var anime repository.Anime
	var username string
	if GetTokenUser(c, &username) {
		var err error
		anime.Anid, err = strconv.Atoi(c.Param("anid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
			return
		}
		if controller.GetSingleAnime(&anime) {
			anime.Status = 2
			controller.UpdateAnime(&anime)
			c.JSON(http.StatusAccepted, gin.H{"status": http.StatusAccepted})
		} else {
			if GetPublishedInfo(&anime) {
				controller.AddNewAnime(&anime)
				c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusAccepted})
			}
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden})
	}
}

type Aidb struct {
	Official_name string
	Zh_name       string
}

func GetBgmName(c *gin.Context) {
	var aidb Aidb
	resp, err := http.Get("https://snow.hacg.top/atobzh.php?anidb=" + c.Query("aid"))
	json.NewDecoder(resp.Body).Decode(&aidb)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
	} else {
		if resp.StatusCode == http.StatusOK {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": aidb})
		} else if resp.StatusCode == http.StatusNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound})
		}
	}
}

func GetAnimeInfo(c *gin.Context) {
	var aidb Aidb
	f, err := os.Open("anime-titles.xml")
	doc, err := xmlquery.Parse(f)
	query := fmt.Sprintf("//anime[@aid='%s']", c.Query("aid"))
	titles := xmlquery.FindOne(doc, query)
	// 优先序列
	if title := titles.SelectElement("//title[@xml:lang='zh-Hans'][@type='official']"); title != nil {
		aidb.Zh_name = title.InnerText()
	} else if title := titles.SelectElement("//title[@xml:lang='zh'][@type='official']"); title != nil {
		aidb.Zh_name = title.InnerText()
	} else if title := titles.SelectElement("//title[@xml:lang='zh-Hans']"); title != nil {
		aidb.Zh_name = title.InnerText()
	} else if title := titles.SelectElement("//title[@xml:lang='zh']"); title != nil {
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

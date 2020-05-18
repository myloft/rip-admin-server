package controller

import (
	"main/repository"
)

func GetAllAnimes(animes *[]repository.Anime, query string, pagenum int, pagesize int, status string) bool {
	if pagenum != 0 && pagesize != 0 {
		Db.Debug().Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("anid = ? OR zh_name like ?", query, "%"+query+"%").Find(animes).RecordNotFound()
	} else {
		Db.Debug().Order("Created_At desc").Where("anid = ? OR zh_name like ?", query, "%"+query+"%").Find(animes).RecordNotFound()
	}
	if Db.Model(animes).RecordNotFound() {
		return false
	}
	return true
}

func GetAllAnimesNum() int {
	var total int
	var animes []repository.Anime
	Db.Find(&animes).Count(&total)
	return total
}

func GetSingleAnime(anid string, anime *repository.Anime) bool {
	if Db.Where("anid = ?", anid).Find(&anime).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func GetUserAnimes(animes *[]repository.Anime, username string, query string, pagenum int, pagesize int, status int) bool {
	if pagenum != 0 && pagesize != 0 {
		Db.Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("owner = ? AND (anid = ? OR zh_name like ?)", username, query, "%"+query+"%").Find(animes).RecordNotFound()
	} else {
		if Db.Order("Created_At desc").Where("owner = ? AND (anid = ? OR zh_name like ?)", username, query, "%"+query+"%").Find(animes).RecordNotFound() {
		}
		return false
	}
	if Db.Model(animes).RecordNotFound() {
		return false
	}
	return true
}

func GetUserAnimesNum(username string) int {
	var total int
	var animes []repository.Anime
	Db.Where("Owner = ?", username).Find(&animes).Count(&total)
	return total
}

func AddNewAnime(anime *repository.Anime) {
	//var anime repository.Anime
	//anime.Anid, _ = strconv.Atoi(c.PostForm("Anid"))
	//anime.Owner = c.PostForm("Owner")
	//anime.Official_name = c.PostForm("Official_name")
	//anime.Zh_name = c.PostForm("Zh_name")
	//anime.Status, _ = strconv.Atoi(c.PostForm("Status"))
	Db.Create(&anime)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": anime})
}

func UpdateAnime(anime *repository.Anime) {
	Db.Model(&anime).Update(&anime)
}

//func UpdateAnime(c *gin.Context) {
//	var anime repository.Anime
//	anime.Anid, _ = strconv.Atoi(c.PostForm("Anid"))
//	anime.Owner = c.PostForm("Owner")
//	anime.Official_name = c.PostForm("Official_name")
//	anime.Zh_name = c.PostForm("Zh_name")
//	anime.Status, _ = strconv.Atoi(c.PostForm("Status"))
//	Db.Model(&anime).Updates(anime)
//	Db.Where("anid = ?", c.Param("anid")).Find(&anime)
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": anime})
//}

func DeleteAnime(anime *repository.Anime) {
	Db.Delete(anime)
}

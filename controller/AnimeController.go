package controller

import (
	"main/repository"
)

func GetAllAnimes(animes *[]repository.Anime, query string, pagenum int, pagesize int, status int) bool {
	var total int
	if status != 0 {
		if pagenum != 0 && pagesize != 0 {
			Db.Debug().Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("anid = ? OR zh_name like ? AND status = ?", query, "%"+query+"%", status).Find(animes).Count(&total).RecordNotFound()
		} else {
			Db.Debug().Order("Created_At desc").Where("anid = ? OR zh_name like ? status = ?", query, "%"+query+"%", status).Find(animes).Count(&total).RecordNotFound()
		}
	} else {
		if pagenum != 0 && pagesize != 0 {
			Db.Debug().Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("anid = ? OR zh_name like ?", query, "%"+query+"%").Find(animes).Count(&total).RecordNotFound()
		} else {
			Db.Debug().Order("Created_At desc").Where("anid = ? OR zh_name like ?", query, "%"+query+"%").Find(animes).Count(&total).RecordNotFound()
		}
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

func GetSingleAnime(anime *repository.Anime) bool {
	if Db.Where("anid = ?", anime.Anid).Find(&anime).RecordNotFound() {
		return false
	} else {
		return true
	}
}

func GetUserAnimes(animes *[]repository.Anime, username string, query string, pagenum int, pagesize int, status int) bool {
	if IsAdmin(username) {
		username = "*"
	}
	if status != 0 {
		if pagenum != 0 && pagesize != 0 {
			Db.Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("owner = ? AND (anid = ? OR zh_name like ?) AND status = ?", username, query, "%"+query+"%", status).Find(animes).RecordNotFound()
		} else {
			if Db.Order("Created_At desc").Where("owner = ? AND (anid = ? OR zh_name like ?) AND status = ?", username, query, "%"+query+"%", status).Find(animes).RecordNotFound() {
			}
			return false
		}
	} else {
		if pagenum != 0 && pagesize != 0 {
			Db.Order("Created_At desc").Limit(pagesize).Offset((pagenum-1)*pagesize).Where("owner = ? AND (anid = ? OR zh_name like ?)", username, query, "%"+query+"%").Find(animes).RecordNotFound()
		} else {
			if Db.Order("Created_At desc").Where("owner = ? AND (anid = ? OR zh_name like ?)", username, query, "%"+query+"%").Find(animes).RecordNotFound() {
			}
			return false
		}
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
	Db.Create(&anime)
}

func UpdateAnime(anime *repository.Anime) {
	Db.Model(&anime).Update(&anime)
}

func DeleteAnime(anime *repository.Anime) {
	Db.Delete(anime)
}

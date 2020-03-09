package controller

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"main/repository"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("sqlite3", "anime.db")
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&repository.User{})
	Db.AutoMigrate(&repository.Anime{})
	Db.AutoMigrate(&repository.Menu{})
}
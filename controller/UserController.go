package controller

import (
	"main/repository"
)

func HavingUser(username string) bool {
	var user repository.User
	Db.Where("name = ?", username).Find(&user).RecordNotFound()
	if Db.Model(&user).RecordNotFound() {
		return false
	}
	return true
}

func GetPassword(username string) string {
	var user repository.User
	Db.Where("name = ?", username).Find(&user)
	return user.Password
}

func GetAllUsers() []repository.User {
	var users []repository.User
	Db.Select("name, admin").Where("name <> ?", "admin").Find(&users)
	return users
}

func UpdateUser(user repository.User) bool {
	err := Db.Model(&user).Updates(user).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

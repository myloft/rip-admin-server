package controller

import "main/repository"

func GetChildrenMenus(menu *repository.Menu) bool{
	if 	Db.Where("Parentid = ?", menu.Id).Find(&menu.Children).RecordNotFound() {
		return false
	} else {
		return true
	}
}
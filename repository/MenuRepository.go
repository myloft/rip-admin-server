package repository

type Menu struct {
	Id       int    `gorm:"PRIMARY_KEY; NOT NULL; UNIQUE_INDEX"`
	Name     string `gorm:"NOT NULL; INDEX"`
	Path     string
	Parentid int    `TYPE:integer REFERENCES Id`
	Children []Menu
}

package repository

type User struct {
	Name     string `gorm:"PRIMARY_KEY; NOT NULL; UNIQUE_INDEX"`
	Admin    *bool
	Password string  `gorm:"NOT NULL"`
	Animes   []Anime `gorm:"foreignkey:Name;association_foreignkey:Owner"`
}

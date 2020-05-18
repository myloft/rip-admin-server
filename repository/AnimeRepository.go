package repository

import "time"

type Anime struct {
	Anid          int    `gorm:"PRIMARY_KEY; NOT NULL; UNIQUE_INDEX"`
	Owner         string `gorm:"NOT NULL; INDEX"`
	Official_name string `gorm:"NOT NULL; UNIQUE_INDEX`
	Zh_name       string `gorm:"NOT NULL; UNIQUE_INDEX`
	Status        int    `gorm:"NOT NULL; INDEX;default: 1`
	CreatedAt     time.Time
}

package dto

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	ID   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(255);"`
}

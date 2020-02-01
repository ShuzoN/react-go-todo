package dto

type User struct {
	ID   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(255);"`
}

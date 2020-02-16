package dto

import "time"

type Todo struct {
	ID       int       `gorm:"AUTO_INCREMENT"`
	Title    string    `gorm:"type:varchar(255)"`
	Deadline time.Time `gorm:"type:datetime"`
	Checked  bool      `gorm:"type:tinyint"`
}

package infrastructures

import (
	"headphonista/src/dto"

	"github.com/jinzhu/gorm"
)

type UserRepositoryOnMysql struct {
	DbConnection *gorm.DB
}

func (userRepository *UserRepositoryOnMysql) GetByID(id int) (*dto.User, error) {

	user := dto.User{}
	if err := userRepository.DbConnection.First(&user, id).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

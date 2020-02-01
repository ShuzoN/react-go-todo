package infrastructures

import (
	"headphonista/src/dto"

	"github.com/jinzhu/gorm"
)

func CreateUserRepositoryOnMysql(connection *gorm.DB) Repository {
	return &UserRepositoryOnMysql{
		dbConnection: connection,
	}
}

type UserRepositoryOnMysql struct {
	dbConnection *gorm.DB
}

func (userRepository *UserRepositoryOnMysql) GetByID(id int) (*dto.User, error) {

	user := dto.User{}
	if err := userRepository.dbConnection.First(&user, id).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

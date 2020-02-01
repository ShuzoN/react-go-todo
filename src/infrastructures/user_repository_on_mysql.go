package infrastructures

import (
	"database/sql"

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

func (userRepository *UserRepositoryOnMysql) GetByID(id int) *sql.Row {
	return userRepository.dbConnection.Raw("select p.id, p.name from users as p where p.id = ?;", id).Row()
}

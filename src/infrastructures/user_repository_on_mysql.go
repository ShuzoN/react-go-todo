package infrastructures

import (
	"database/sql"
)

func CreateUserRepositoryOnMysql(connection *sql.DB) Repository {
	return &UserRepositoryOnMysql{
		dbConnection: connection,
	}
}

type UserRepositoryOnMysql struct {
	dbConnection *sql.DB
}

func (userRepository *UserRepositoryOnMysql) GetByID(id int) *sql.Row {
	return userRepository.dbConnection.QueryRow("select p.id, p.name from users as p where p.id = ?;", id)
}

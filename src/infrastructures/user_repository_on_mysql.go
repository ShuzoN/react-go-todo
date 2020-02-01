package infrastructures

import (
	"database/sql"
	"headphonista/src/bootstrap"
)

func CreateUserRepositoryOnMysql() UserRepository {
	ur := new(UserRepositoryOnMysql)
	ur.dbConnection = bootstrap.Connection.DbConnection
	return ur
}

type UserRepositoryOnMysql struct {
	dbConnection *sql.DB
}

func (userRepository *UserRepositoryOnMysql) GetUserByID(id int) *sql.Row {
	return userRepository.dbConnection.QueryRow("select p.name from users as p where p.id = ?;", id)
}

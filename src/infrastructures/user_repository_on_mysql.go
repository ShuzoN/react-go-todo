package infrastructures

import (
	"database/sql"
	"headphonista/src/bootstrap"
)

func CreateUserRepositoryOnMysql() UserRepository {
	ur := new(UserRepositoryOnMysql)
	ur.bootstrap = bootstrap.Connection
	return ur
}

type UserRepositoryOnMysql struct {
	bootstrap *bootstrap.Bootstrap
}

func (userRepository *UserRepositoryOnMysql) GetUserByID(id int) *sql.Row {
	return userRepository.bootstrap.DbConnection.QueryRow("select p.name from users as p where p.id = ?;", id)
}

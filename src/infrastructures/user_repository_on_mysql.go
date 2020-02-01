package infrastructures

import (
	"database/sql"
	"headphonista/src/bootstrap"
)

// UserRepositoryOnMysql はuserRepositoryの実装クラス
type UserRepositoryOnMysql struct {
	Bootstrap *bootstrap.Bootstrap
}

func (userRepository *UserRepositoryOnMysql) GetUserByID(id int) *sql.Row {
	return userRepository.Bootstrap.DbConnection.QueryRow("select p.name from users as p where p.id = ?;", id)
}

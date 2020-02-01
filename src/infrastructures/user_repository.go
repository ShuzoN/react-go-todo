package infrastructures

import "database/sql"

type UserRepository interface {
	GetUserByID(id int) *sql.Row
}

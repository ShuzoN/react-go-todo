package infrastructures

import "database/sql"

type Repository interface {
	GetByID(id int) *sql.Row
}

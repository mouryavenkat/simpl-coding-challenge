package mysql

import (
	"database/sql"
)

type mysqlWrapper struct {
	dbObject *sql.DB
}

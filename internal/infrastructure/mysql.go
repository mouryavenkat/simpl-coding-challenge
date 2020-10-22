package infrastructure

import "database/sql"

type MySQL interface {
	Insert(query string, args ...interface{}) error
	Update(query string, args ...interface{}) error
	Query(query string, args ...interface{}) (Rows, error)
	Begin() (*sql.Tx, error)
}

type Rows interface {
	Close() error
	Scan(dest ...interface{}) error
	Next() bool
}

type DB interface {
	Rollback() error
	Commit() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

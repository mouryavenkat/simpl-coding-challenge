package mysql

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"simpl-coding-challenge/internal/infrastructure"
)

func NewSqlWrapper(db *sql.DB) infrastructure.MySQL {
	return &mysqlWrapper{
		dbObject: db,
	}
}

func (m *mysqlWrapper) Insert(query string, args ...interface{}) error {
	statement, err := m.dbObject.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(args...)
	if err != nil {
		if errorType, ok := err.(*mysql.MySQLError); ok {
			if errorType.Number == 1062 {
				return errors.New("can't create a duplicate name")
			}
		}
		return err
	}
	return nil
}

func (m *mysqlWrapper) Query(query string, args ...interface{}) (rows infrastructure.Rows, err error) {
	rows, err = m.dbObject.Query(query, args...)
	return
}

func (m *mysqlWrapper) Update(query string, args ...interface{}) error {
	statement, err := m.dbObject.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlWrapper) Begin() (*sql.Tx, error) {
	txn, _ := m.dbObject.Begin()
	return txn, nil
}

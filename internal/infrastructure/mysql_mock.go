package infrastructure

import (
	"database/sql"
)

type MockMySQL interface {
	MySQL
	WithInsertError(err error) MockMySQL
	WithUpdateError(err error) MockMySQL
	WithQueryError(err error) MockMySQL
	WithQueryRows(rows Rows) MockMySQL
}

type mockMySQL struct {
	insertError, updateError, queryError error
	queryCount                           int
	queryRows                            Rows
}

func (m mockMySQL) Insert(query string, args ...interface{}) error {
	return m.insertError
}

func (m mockMySQL) WithInsertError(err error) MockMySQL {
	m.insertError = err
	return m
}

func (m mockMySQL) Update(query string, args ...interface{}) error {
	return m.updateError
}

func (m mockMySQL) WithUpdateError(err error) MockMySQL {
	m.updateError = err
	return m
}

func (m mockMySQL) Query(query string, args ...interface{}) (Rows, error) {
	return m.queryRows, m.queryError
}

func (m mockMySQL) WithQueryRows(rows Rows) MockMySQL {
	m.queryRows = rows
	return m
}

func (m mockMySQL) WithQueryError(err error) MockMySQL {
	m.queryError = err
	return m
}

func (m mockMySQL) Begin() (*sql.Tx, error) {
	panic("implement me")
}

func NewMockMySQL() MockMySQL {
	return &mockMySQL{}
}

type MockRows interface {
	Rows
	WithRecordCount(count int) MockRows
	WithScanResult(result [][]interface{}) MockRows
}

type mockRows struct {
	recordCount, nextCount int
	closeError             error
	scanResult             [][]interface{}
	scanCount              int
}

func (m *mockRows) Close() error {
	return m.closeError
}

func (m *mockRows) Scan(dest ...interface{}) error {
	dest = []interface{}{dest}
	for index := range dest {
		var x interface{} = m.scanResult[m.scanCount][index]
		dest[index] = &x
	}
	m.scanCount = m.scanCount + 1
	return nil
}

func (m *mockRows) WithScanResult(result [][]interface{}) MockRows {
	m.scanResult = result
	return m
}

func (m *mockRows) WithRecordCount(count int) MockRows {
	m.recordCount = count
	return m
}

func (m *mockRows) Next() bool {
	if m.recordCount == m.nextCount {
		return false
	}
	m.nextCount = m.nextCount + 1
	return true
}

func NewMockRows() MockRows {
	return &mockRows{}
}

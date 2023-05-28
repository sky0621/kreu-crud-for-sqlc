package parser

import (
	"strings"
)

// CreateInitialSQLParseResult is
func CreateInitialSQLParseResult(sqlName, sqlFileName string) *SQLParseResult {
	return &SQLParseResult{SQLName: ToSQLName(sqlName), SQLFileName: ToSQLFileName(sqlFileName)}
}

// SQLParseResult is
type SQLParseResult struct {
	SQLName                SQLName
	SQLFileName            SQLFileName
	TableNameWithCRUDSlice []*TableNameWithCRUD
}

func createTableNameWithCRUD(tableName string, crud CRUD) *TableNameWithCRUD {
	return &TableNameWithCRUD{TableName: TableName(tableName), CRUD: crud}
}

type TableNameWithCRUD struct {
	TableName TableName
	CRUD      CRUD
}

type TableName string

func (t TableName) ToString() string {
	return string(t)
}

type SQLName string

func (n SQLName) ToString() string {
	return string(n)
}

func ToSQLName(n string) SQLName {
	return SQLName(strings.Trim(n, " "))
}

type SQLFileName string

func (n SQLFileName) ToString() string {
	return string(n)
}

func ToSQLFileName(n string) SQLFileName {
	return SQLFileName(strings.Trim(n, " "))
}

package parser

import (
	"strings"

	query "github.com/pganalyze/pg_query_go/v4"
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

func parseNodes(nodes []*query.Node, crud CRUD, results []*TableNameWithCRUD) []*TableNameWithCRUD {
	for _, node := range nodes {
		results = append(results, ParseNode(node, crud)...)
	}
	return results
}

func parseNodesNodes(nodesNodes [][]*query.Node, crud CRUD, results []*TableNameWithCRUD) []*TableNameWithCRUD {
	for _, nodes := range nodesNodes {
		for _, node := range nodes {
			results = append(results, ParseNode(node, crud)...)
		}
	}
	return results
}

package internal

import (
	"errors"

	"github.com/sky0621/kreu-crud-for-sqlc/internal/parser"

	query "github.com/pganalyze/pg_query_go/v4"
)

type SQLParser interface {
	Parse(sqlName, sqlFileName, sql string) (*parser.SQLParseResult, error)
}

func NewSQLParser() SQLParser {
	return &sqlParser{}
}

type sqlParser struct {
	// Config ?
}

func (p *sqlParser) Parse(sqlName, sqlFileName, sql string) (*parser.SQLParseResult, error) {
	res, err := query.Parse(sql)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("result == nil")
	}

	processedResult := parser.CreateInitialSQLParseResult(sqlName, sqlFileName)

	for _, stmt := range res.GetStmts() {
		tableNameWithCRUDSlice := parser.ParseNode(stmt.GetStmt(), parser.Undecided)
		if tableNameWithCRUDSlice == nil || len(tableNameWithCRUDSlice) == 0 {
			continue
		}
		processedResult.TableNameWithCRUDSlice = append(processedResult.TableNameWithCRUDSlice, tableNameWithCRUDSlice...)
	}

	return processedResult, nil
}

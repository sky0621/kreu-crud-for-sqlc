package internal

import (
	"errors"

	query "github.com/pganalyze/pg_query_go/v4"
	"github.com/sky0621/kreu-crud-for-sqlc/internal/parser"
)

type sqlParser2 struct {
	// Config ?
}

func (p *sqlParser2) Parse(sqlName, sqlFileName, sql string) (*parser.SQLParseResult, error) {
	res, err := query.Parse(sql)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("result == nil")
	}

	result := parser.CreateInitialSQLParseResult(sqlName, sqlFileName)

	for _, stmt := range res.GetStmts() {
		//		tableNameWithCRUDs := parser.ExamineNode(stmt.GetStmt(), parser.Undecided)
		//	tableNameWithCRUDs := parser.ExamineNode(res, parser.Undecided)
		//		if tableNameWithCRUDs == nil || len(tableNameWithCRUDs) == 0 {
		//			continue
		//	return result, nil
		//		}
		//		result.TableNameWithCRUDSlice = append(result.TableNameWithCRUDSlice, tableNameWithCRUDs...)
		tableNameWithCRUDs := parser.ExamineTables(stmt.GetStmt(), parser.Undecided)
		if tableNameWithCRUDs == nil || len(tableNameWithCRUDs) == 0 {
			continue
		}
		result.TableNameWithCRUDSlice = append(result.TableNameWithCRUDSlice, tableNameWithCRUDs...)
	}

	return result, nil
}

package internal

import (
	"errors"
	"log"
	"strings"

	query "github.com/pganalyze/pg_query_go/v4"
)

type SQLParser interface {
	Parse(sqlName, sqlFileName, sql string) (*SQLParseResult, error)
}

func NewSQLParser() SQLParser {
	return &sqlParser{}
}

type sqlParser struct {
}

func (p *sqlParser) Parse(sqlName, sqlFileName, sql string) (*SQLParseResult, error) {
	res, err := query.Parse(sql)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("result == nil")
	}

	processedResult := &SQLParseResult{SQLName: ToSQLName(sqlName), SQLFileName: ToSQLFileName(sqlFileName)}

	for _, stmt := range res.GetStmts() {
		tableNameWithCRUDSlice, err := parseStmt(stmt.GetStmt())
		if err != nil {
			log.Println("parseStmt is nil")
			continue
		}
		if tableNameWithCRUDSlice == nil || len(tableNameWithCRUDSlice) == 0 {
			continue
		}
		processedResult.TableNameWithCRUDSlice = append(processedResult.TableNameWithCRUDSlice, tableNameWithCRUDSlice...)
	}

	return processedResult, nil
}

type SQLParseResult struct {
	SQLName                SQLName
	SQLFileName            SQLFileName
	TableNameWithCRUDSlice []*TableNameWithCRUD
}

type TableNameWithCRUD struct {
	TableName TableName
	CRUD      CRUD
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

type TableName string

func (t TableName) ToString() string {
	return string(t)
}

func parseStmt(s *query.Node) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, errors.New("node is nil")
	}

	var tableNameWithCRUDSlice []*TableNameWithCRUD
	var err error

	tableNameWithCRUDSlice = parseSelectStmt(s.GetSelectStmt())
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseInsertStmt(s.GetInsertStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseUpdateStmt(s.GetUpdateStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	tableNameWithCRUDSlice, err = parseDeleteStmt(s.GetDeleteStmt())
	if err != nil {
		return nil, err
	}
	if tableNameWithCRUDSlice != nil {
		return tableNameWithCRUDSlice, nil
	}

	return nil, nil
}

func parseNodeRangeVar(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return result
	}

	nv, ok := n.(*query.Node_RangeVar)
	if !ok {
		return result
	}

	if nv == nil {
		return result
	}

	if nv.RangeVar == nil {
		return result
	}

	result = append(result, &TableNameWithCRUD{CRUD: crud, TableName: TableName(nv.RangeVar.Relname)})
	return result
}

func parseNodeJoinExpr(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return result
	}

	nj, ok := n.(*query.Node_JoinExpr)
	if !ok {
		return result
	}

	if nj == nil {
		return result
	}

	if nj.JoinExpr == nil {
		return result
	}

	if nj.JoinExpr.Larg != nil {
		res := parseNode(nj.JoinExpr.Larg, crud)
		if res != nil {
			result = append(result, res...)
		}
	}

	if nj.JoinExpr.Rarg != nil {
		res := parseNode(nj.JoinExpr.Rarg, crud)
		if res != nil {
			result = append(result, res...)
		}
	}

	return result
}

func parseNodeRangeSubSelect(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return result
	}

	rs, ok := n.(*query.Node_RangeSubselect)
	if !ok {
		return result
	}

	if rs == nil {
		return result
	}

	if rs.RangeSubselect == nil {
		return result
	}

	sq := rs.RangeSubselect.GetSubquery()
	if sq == nil {
		return result
	}

	sRes := parseSelectStmt(sq.GetSelectStmt())
	if sRes != nil {
		result = append(result, sRes...)
	}

	return result
}

func parseNode(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD

	res := parseNodeRangeVar(node, crud)
	if res != nil {
		result = append(result, res...)
	}

	res2 := parseNodeJoinExpr(node, crud)
	if res2 != nil {
		result = append(result, res2...)
	}

	res3 := parseNodeRangeSubSelect(node, crud)
	if res3 != nil {
		result = append(result, res3...)
	}

	return result
}

func parseSelectStmt(s *query.SelectStmt) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD

	if s == nil {
		return result
	}

	crud := Read

	for _, from := range s.FromClause {
		res := parseNode(from, crud)
		if res != nil {
			result = append(result, res...)
		}
	}

	lRes := parseSelectStmt(s.Larg)
	if lRes != nil {
		result = append(result, lRes...)
	}

	rRes := parseSelectStmt(s.Rarg)
	if rRes != nil {
		result = append(result, rRes...)
	}

	return result
}

func parseInsertStmt(s *query.InsertStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Create

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseUpdateStmt(s *query.UpdateStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Update

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

func parseDeleteStmt(s *query.DeleteStmt) ([]*TableNameWithCRUD, error) {
	if s == nil {
		return nil, nil
	}

	var result []*TableNameWithCRUD
	crud := Delete

	rel := s.GetRelation()
	if rel != nil {
		result = append(result, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return result, nil
}

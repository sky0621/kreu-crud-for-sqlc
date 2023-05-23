package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseSelectStmt(s *query.SelectStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Read

	for _, from := range s.FromClause {
		results = append(results, ParseNode(from, crud)...)
	}

	results = append(results, parseSelectStmt(s.Larg)...)
	results = append(results, parseSelectStmt(s.Rarg)...)

	wh := s.WhereClause
	if wh != nil {
		whRes := parseSelectStmt(wh.GetSelectStmt())
		if whRes != nil {
			results = append(results, whRes...)
		}
	}

	return results
}

func parseInsertStmt(s *query.InsertStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Create

	rel := s.GetRelation()
	if rel != nil {
		results = append(results, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return results
}

func parseUpdateStmt(s *query.UpdateStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Update

	rel := s.GetRelation()
	if rel != nil {
		results = append(results, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return results
}

func parseDeleteStmt(s *query.DeleteStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Delete

	rel := s.GetRelation()
	if rel != nil {
		results = append(results, &TableNameWithCRUD{CRUD: crud, TableName: TableName(rel.GetRelname())})
	}

	return results
}

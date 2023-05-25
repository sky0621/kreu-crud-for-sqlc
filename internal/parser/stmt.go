package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseSelectStmt(s *query.SelectStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Read

	results = parseNodes([]*query.Node{
		s.GetHavingClause(),
		s.GetLimitCount(),
		s.GetLimitOffset(),
		s.GetWhereClause(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		s.GetDistinctClause(),
		s.GetFromClause(),
		s.GetGroupClause(),
		s.GetLockingClause(),
		s.GetSortClause(),
		s.GetTargetList(),
		s.GetValuesLists(),
		s.GetWindowClause(),
	}, crud, results)

	selectStmts := []*query.SelectStmt{
		s.GetLarg(),
		s.GetRarg(),
	}
	for _, selectStmt := range selectStmts {
		results = append(results, parseSelectStmt(selectStmt)...)
	}

	results = append(results, parseIntoClause(s.GetIntoClause(), crud)...)
	results = append(results, parseWithClause(s.GetWithClause(), crud)...)

	return results
}

func parseInsertStmt(s *query.InsertStmt) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if s == nil {
		return results
	}

	crud := Create

	results = parseNodes([]*query.Node{
		s.GetSelectStmt(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		s.GetCols(),
		s.GetReturningList(),
	}, crud, results)

	results = append(results, parseWithClause(s.GetWithClause(), crud)...)
	results = append(results, parseOnConflictClause(s.GetOnConflictClause(), crud)...)

	results = append(results, parseRangeVar(s.GetRelation(), crud)...)

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

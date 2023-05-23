package parser

import query "github.com/pganalyze/pg_query_go/v4"

func ParseNode(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil || node.GetNode() == nil {
		return results
	}

	results = append(results, parseSelectStmt(node.GetSelectStmt())...)
	results = append(results, parseInsertStmt(node.GetInsertStmt())...)
	results = append(results, parseUpdateStmt(node.GetUpdateStmt())...)
	results = append(results, parseDeleteStmt(node.GetDeleteStmt())...)

	rv := node.GetRangeVar()
	if rv != nil {
		results = append(results, createTableNameWithCRUD(rv.Relname, crud))
	}

	n := node.GetNode()

	nv, ok := n.(*query.Node_RangeVar)
	if ok && nv != nil && nv.RangeVar != nil {
		results = append(results, createTableNameWithCRUD(nv.RangeVar.Relname, crud))
		return results
	}

	nj, ok := n.(*query.Node_JoinExpr)
	if ok && nj != nil && nj.JoinExpr != nil {
		nj.JoinExpr.GetLarg()
	}

	return results
}

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

func parseNodeRangeVar(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return results
	}

	nv, ok := n.(*query.Node_RangeVar)
	if !ok {
		return results
	}

	if nv == nil {
		return results
	}

	if nv.RangeVar == nil {
		return results
	}

	results = append(results, &TableNameWithCRUD{CRUD: crud, TableName: TableName(nv.RangeVar.Relname)})
	return results
}

func parseNodeJoinExpr(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return results
	}

	nj, ok := n.(*query.Node_JoinExpr)
	if !ok {
		return results
	}

	if nj == nil {
		return results
	}

	if nj.JoinExpr == nil {
		return results
	}

	if nj.JoinExpr.Larg != nil {
		res := ParseNode(nj.JoinExpr.Larg, crud)
		if res != nil {
			results = append(results, res...)
		}
	}

	if nj.JoinExpr.Rarg != nil {
		res := ParseNode(nj.JoinExpr.Rarg, crud)
		if res != nil {
			results = append(results, res...)
		}
	}

	return results
}

func parseNodeRangeSubSelect(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return results
	}

	rs, ok := n.(*query.Node_RangeSubselect)
	if !ok {
		return results
	}

	if rs == nil {
		return results
	}

	if rs.RangeSubselect == nil {
		return results
	}

	sq := rs.RangeSubselect.GetSubquery()
	if sq == nil {
		return results
	}

	sRes := parseSelectStmt(sq.GetSelectStmt())
	if sRes != nil {
		results = append(results, sRes...)
	}

	return results
}

func parseNodeAExpr(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return results
	}

	na, ok := n.(*query.Node_AExpr)
	if !ok {
		return results
	}

	if na == nil {
		return results
	}

	if na.AExpr == nil {
		return results
	}

	lx := na.AExpr.GetLexpr()
	if lx != nil {
		results = append(results, ParseNode(lx, crud)...)
	}

	rx := na.AExpr.GetRexpr()
	if rx != nil {
		results = append(results, ParseNode(rx, crud)...)
	}

	return results
}

func parseWhereClause(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	n := node.GetNode()
	if n == nil {
		return results
	}

	return results
}

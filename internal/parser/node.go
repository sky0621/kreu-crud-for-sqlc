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

	results = append(results, parseJoinExpr(node.GetJoinExpr(), crud)...)
	results = append(results, parseFromExpr(node.GetFromExpr(), crud)...)

	n := node.GetNode()

	nj, ok := n.(*query.Node_JoinExpr)
	if ok && nj != nil && nj.JoinExpr != nil {
		nj.JoinExpr.GetLarg()
	}

	return results
}

func parseJoinExpr(node *query.JoinExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetLarg(), crud)...)
	results = append(results, ParseNode(node.GetRarg(), crud)...)

	return results
}

func parseFromExpr(node *query.FromExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	// FIXME:

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

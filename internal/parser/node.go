package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseNode(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var result []*TableNameWithCRUD
	if node == nil || node.GetNode() == nil {
		return result
	}

	n := node.GetNode()

	nv, ok := n.(*query.Node_RangeVar)
	if ok && nv != nil && nv.RangeVar != nil {
		result = append(result, &TableNameWithCRUD{CRUD: crud, TableName: TableName(nv.RangeVar.Relname)})
		return result
	}

	nj, ok := n.(*query.Node_JoinExpr)
	if ok && nj != nil && nj.JoinExpr != nil {
		nj.JoinExpr.GetLarg()
	}

	node.GetSelectStmt()
	return result
}

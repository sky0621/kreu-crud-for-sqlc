package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseCaseExpr(node *query.CaseExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetArg(),
		node.GetDefresult(),
		node.GetXpr(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetArgs(),
	}, crud, results)

	return results
}

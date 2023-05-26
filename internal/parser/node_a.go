package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseAccessPriv(node *query.AccessPriv, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodesNodes([][]*query.Node{
		node.GetCols(),
	}, crud, results)

	return results
}

func parseAArrayExpr(node *query.A_ArrayExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodesNodes([][]*query.Node{
		node.GetElements(),
	}, crud, results)

	return results
}

func parseAggref(node *query.Aggref, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetAggfilter(),
		node.GetXpr(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetAggargtypes(),
		node.GetAggdirectargs(),
		node.GetAggdistinct(),
		node.GetAggorder(),
		node.GetArgs(),
	}, crud, results)

	return results
}

func parseAExpr(node *query.A_Expr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetLexpr(),
		node.GetRexpr(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetName(),
	}, crud, results)

	return results
}

func parseAIndices(node *query.A_Indices, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetLidx(),
		node.GetUidx(),
	}, crud, results)

	return results
}

func parseAIndirection(node *query.A_Indirection, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetArg(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetIndirection(),
	}, crud, results)

	return results
}

func parseAlias(node *query.Alias, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodesNodes([][]*query.Node{
		node.GetColnames(),
	}, crud, results)

	return results
}

func parseArrayCoerceExpr(node *query.ArrayCoerceExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetArg(),
		node.GetElemexpr(),
		node.GetXpr(),
	}, crud, results)

	return results
}

func parseArrayExpr(node *query.ArrayExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetXpr(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetElements(),
	}, crud, results)

	return results
}

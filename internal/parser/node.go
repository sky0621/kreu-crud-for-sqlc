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

	results = append(results, parseRangeVar(node.GetRangeVar(), crud)...)

	results = append(results, parseFromExpr(node.GetFromExpr(), crud)...)
	results = append(results, parseJoinExpr(node.GetJoinExpr(), crud)...)
	results = append(results, parseRangeSubSelect(node.GetRangeSubselect(), crud)...)

	results = append(results, parseAccessPriv(node.GetAccessPriv(), crud)...)
	//node.GetAConst()
	results = append(results, parseAArrayExpr(node.GetAArrayExpr(), crud)...)
	results = append(results, parseAggref(node.GetAggref(), crud)...)
	results = append(results, parseAExpr(node.GetAExpr(), crud)...)
	results = append(results, parseAIndices(node.GetAIndices(), crud)...)
	node.GetAIndirection()
	node.GetAlias()
	node.GetAlterCollationStmt()
	node.GetAlterDatabaseRefreshCollStmt()
	node.GetAlterDatabaseSetStmt()
	node.GetAlterDatabaseStmt()
	node.GetAlterDefaultPrivilegesStmt()
	node.GetAlterDomainStmt()
	node.GetAlterEnumStmt()
	node.GetAlterEventTrigStmt()
	node.GetAlterExtensionContentsStmt()
	node.GetAlterExtensionStmt()
	node.GetAlterFdwStmt()
	node.GetAlterForeignServerStmt()
	node.GetAlterFunctionStmt()

	return results
}

func parseRangeVar(v *query.RangeVar, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if v == nil {
		return results
	}

	results = append(results, createTableNameWithCRUD(v.Relname, crud))

	return results
}

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

func parseFromExpr(node *query.FromExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetQuals(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetFromlist(),
	}, crud, results)

	return results
}

func parseJoinExpr(node *query.JoinExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		node.GetLarg(),
		node.GetQuals(),
		node.GetRarg(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		node.GetUsingClause(),
	}, crud, results)

	aliases := []*query.Alias{
		node.GetAlias(),
		node.GetJoinUsingAlias(),
	}
	for _, alias := range aliases {
		results = append(results, parseAlias(alias, crud)...)
	}

	return results
}

func parseAlias(a *query.Alias, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if a == nil {
		return results
	}

	results = parseNodesNodes([][]*query.Node{
		a.GetColnames(),
	}, crud, results)

	return results
}
func parseRangeSubSelect(node *query.RangeSubselect, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetSubquery(), crud)...)

	return results
}

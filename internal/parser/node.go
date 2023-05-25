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

	for _, e := range node.GetCols() {
		results = append(results, ParseNode(e, crud)...)
	}

	return results
}

func parseAArrayExpr(node *query.A_ArrayExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	for _, e := range node.GetElements() {
		results = append(results, ParseNode(e, crud)...)
	}

	return results
}

func parseAggref(node *query.Aggref, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	for _, e := range node.GetAggargtypes() {
		results = append(results, ParseNode(e, crud)...)
	}
	for _, e := range node.GetAggdirectargs() {
		results = append(results, ParseNode(e, crud)...)
	}
	for _, e := range node.GetAggdistinct() {
		results = append(results, ParseNode(e, crud)...)
	}
	for _, e := range node.GetAggorder() {
		results = append(results, ParseNode(e, crud)...)
	}
	for _, e := range node.GetArgs() {
		results = append(results, ParseNode(e, crud)...)
	}
	results = append(results, ParseNode(node.GetAggfilter(), crud)...)
	results = append(results, ParseNode(node.GetXpr(), crud)...)

	return results
}

func parseAExpr(node *query.A_Expr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetLexpr(), crud)...)
	results = append(results, ParseNode(node.GetRexpr(), crud)...)
	for _, e := range node.GetName() {
		results = append(results, ParseNode(e, crud)...)
	}

	return results
}

func parseAIndices(node *query.A_Indices, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetLidx(), crud)...)
	results = append(results, ParseNode(node.GetUidx(), crud)...)

	return results
}

func parseFromExpr(node *query.FromExpr, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	for _, from := range node.GetFromlist() {
		results = append(results, ParseNode(from, crud)...)
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

func parseRangeSubSelect(node *query.RangeSubselect, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetSubquery(), crud)...)

	return results
}

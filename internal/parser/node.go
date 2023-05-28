package parser

import (
	query "github.com/pganalyze/pg_query_go/v4"
)

func ExamineNode(n interface{}, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if n == nil {
		return results
	}

	node := toNode(n)
	if node == nil {
		return results
	}

	useCRUD := crud

	actualCRUD := JudgeCRUD(node)
	if actualCRUD != Undecided {
		useCRUD = actualCRUD
	}

	results = append(results, ExamineTables(n, useCRUD)...)

	return results
}

func toNode(v any) *query.Node {
	node, ok := v.(*query.Node)
	if !ok {
		return nil
	}
	return node
}

func getTableName(v *query.RangeVar) string {
	if v == nil {
		return ""
	}
	return v.GetRelname()
}

// 消す？
func ParseNode(node *query.Node, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	/*
	 * parse statement
	 */
	results = append(results, parseSelectStmt(node.GetSelectStmt())...)
	results = append(results, parseInsertStmt(node.GetInsertStmt())...)
	results = append(results, parseUpdateStmt(node.GetUpdateStmt())...)
	results = append(results, parseDeleteStmt(node.GetDeleteStmt())...)
	results = append(results, parseAlterCollationStmt(node.GetAlterCollationStmt(), crud)...)

	/*
		{
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
			node.GetAlternativeSubPlan()
			node.GetAlterObjectDependsStmt()
			node.GetAlterObjectSchemaStmt()
			node.GetAlterOperatorStmt()
			node.GetAlterOpFamilyStmt()
			node.GetAlterOwnerStmt()
			node.GetAlterPolicyStmt()
			node.GetAlterPublicationStmt()
			node.GetAlterRoleSetStmt()
			node.GetAlterRoleStmt()
			node.GetAlterSeqStmt()
			node.GetAlterStatsStmt()
			node.GetAlterSubscriptionStmt()
			node.GetAlterSystemStmt()
			node.GetAlterTableCmd()
			node.GetAlterTableMoveAllStmt()
			node.GetAlterTableSpaceOptionsStmt()
			node.GetAlterTableStmt()
			node.GetAlterTsconfigurationStmt()
			node.GetAlterTsdictionaryStmt()
			node.GetAlterTypeStmt()
			node.GetAlterUserMappingStmt()
		}
	*/

	node.GetCallStmt()

	/*
	 * pickup table name
	 */
	results = append(results, parseRangeVar(node.GetRangeVar(), crud)...)

	/*
	 * parse others
	 */
	results = append(results, parseAccessPriv(node.GetAccessPriv(), crud)...)
	//node.GetAConst()
	results = append(results, parseAArrayExpr(node.GetAArrayExpr(), crud)...)
	results = append(results, parseAggref(node.GetAggref(), crud)...)
	results = append(results, parseAExpr(node.GetAExpr(), crud)...)
	results = append(results, parseAIndices(node.GetAIndices(), crud)...)
	results = append(results, parseAIndirection(node.GetAIndirection(), crud)...)
	results = append(results, parseAlias(node.GetAlias(), crud)...)
	results = append(results, parseArrayCoerceExpr(node.GetArrayCoerceExpr(), crud)...)
	results = append(results, parseArrayExpr(node.GetArrayExpr(), crud)...)
	//node.GetAStar()

	//node.GetBitString()
	//node.GetBoolean()
	node.GetBooleanTest()
	node.GetBoolExpr()

	//node.GetCallContext()
	results = append(results, parseCaseExpr(node.GetCaseExpr(), crud)...)
	node.GetCaseTestExpr()
	node.GetCaseWhen()
	node.GetCheckPointStmt()
	node.GetClosePortalStmt()
	node.GetClusterStmt()
	node.GetCoalesceExpr()
	node.GetCoerceToDomain()
	node.GetCoerceToDomainValue()
	node.GetCoerceViaIo()
	node.GetCollateClause()
	node.GetCollateExpr()
	node.GetColumnDef()
	node.GetColumnRef()
	node.GetCommentStmt()
	node.GetCommonTableExpr()
	node.GetCompositeTypeStmt()
	node.GetConstraint()
	node.GetConstraintsSetStmt()
	node.GetConvertRowtypeExpr()
	node.GetCopyStmt()
	node.GetCreateAmStmt()
	node.GetCreateCastStmt()
	node.GetCreateConversionStmt()
	node.GetCreatedbStmt()
	node.GetCreateDomainStmt()
	node.GetCreateEnumStmt()
	node.GetCreateEventTrigStmt()
	node.GetCreateExtensionStmt()
	node.GetCreateFdwStmt()
	node.GetCreateForeignServerStmt()
	node.GetCreateForeignTableStmt()
	// FIXME: GetC~~

	node.GetDeallocateStmt()
	node.GetDeclareCursorStmt()
	node.GetDefElem()
	node.GetDefineStmt()
	node.GetDiscardStmt()
	node.GetDistinctExpr()
	node.GetDoStmt()
	node.GetDropdbStmt()
	node.GetDropSubscriptionStmt()
	node.GetDropOwnedStmt()
	node.GetDropRoleStmt()
	node.GetDropStmt()
	node.GetDropTableSpaceStmt()
	node.GetDropUserMappingStmt()

	results = append(results, parseFromExpr(node.GetFromExpr(), crud)...)

	results = append(results, parseJoinExpr(node.GetJoinExpr(), crud)...)

	results = append(results, parseRangeSubSelect(node.GetRangeSubselect(), crud)...)

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

func parseRangeSubSelect(node *query.RangeSubselect, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if node == nil {
		return results
	}

	results = append(results, ParseNode(node.GetSubquery(), crud)...)

	return results
}

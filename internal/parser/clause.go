package parser

import query "github.com/pganalyze/pg_query_go/v4"

func parseIntoClause(c *query.IntoClause, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if c == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		c.GetViewQuery(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		c.GetColNames(),
		c.GetOptions(),
	}, crud, results)

	results = append(results, parseRangeVar(c.GetRel(), crud)...)

	return results
}

func parseWithClause(c *query.WithClause, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if c == nil {
		return results
	}

	results = parseNodesNodes([][]*query.Node{
		c.GetCtes(),
	}, crud, results)

	return results
}

func parseOnConflictClause(c *query.OnConflictClause, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if c == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		c.GetWhereClause(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		c.GetTargetList(),
	}, crud, results)

	results = append(results, parseInferClause(c.GetInfer(), crud)...)

	return results
}

func parseInferClause(c *query.InferClause, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if c == nil {
		return results
	}

	results = parseNodes([]*query.Node{
		c.GetWhereClause(),
	}, crud, results)

	results = parseNodesNodes([][]*query.Node{
		c.GetIndexElems(),
	}, crud, results)

	return results
}

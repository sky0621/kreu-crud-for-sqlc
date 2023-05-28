package parser

import query "github.com/pganalyze/pg_query_go/v4"

type CRUD int8

const (
	Undecided CRUD = iota
	Create
	Read
	Update
	Delete
)

func (c CRUD) ToName() string {
	switch c {
	case Create:
		return "CREATE"
	case Read:
		return "READ"
	case Update:
		return "UPDATE"
	case Delete:
		return "DELETE"
	}
	return ""
}

func (c CRUD) ToShortName() string {
	switch c {
	case Create:
		return "C"
	case Read:
		return "R"
	case Update:
		return "U"
	case Delete:
		return "D"
	}
	return ""
}

func JudgeCRUD(node *query.Node) CRUD {
	if node == nil {
		return Undecided
	}

	if node.GetSelectStmt() != nil {
		return Read
	}
	if node.GetInsertStmt() != nil {
		return Create
	}
	if node.GetUpdateStmt() != nil {
		return Update
	}
	if node.GetDeleteStmt() != nil {
		return Delete
	}
	return Undecided
}

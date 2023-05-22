package parser

type TableNameWithCRUD struct {
	TableName TableName
	CRUD      CRUD
}

type TableName string

func (t TableName) ToString() string {
	return string(t)
}

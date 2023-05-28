package parser

import (
	"reflect"

	query "github.com/pganalyze/pg_query_go/v4"
)

func ExamineTables(value interface{}, crud CRUD) []*TableNameWithCRUD {
	var results []*TableNameWithCRUD

	if value == nil {
		return results
	}

	useCRUD := crud

	node, ok := value.(*query.Node)
	if ok {
		actualCRUD := judgeCRUD(node)
		if actualCRUD != Undecided {
			useCRUD = actualCRUD
		}
	}

	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if v.Type() == reflect.TypeOf(query.RangeVar{}) {
		rangeVar := value.(query.RangeVar)
		results = append(results, createTableNameWithCRUD(rangeVar.GetRelname(), useCRUD))
	}

	switch t.Kind() {
	case reflect.Ptr:
		if v.Elem().IsValid() {
			results = append(results, ExamineTables(v.Elem().Interface(), useCRUD)...)
		}
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		if v.Len() > 0 {
			for i := 0; i < v.Len(); i++ {
				results = append(results, ExamineTables(v.Index(i).Interface(), useCRUD)...)
			}
		}
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			switch f.Type.String() {
			case "impl.MessageState":
				fallthrough
			case "int32":
				fallthrough
			case "[]uint8":
			default:
				results = append(results, ExamineTables(reflect.ValueOf(value).Field(i).Interface(), useCRUD)...)
			}
		}
	}
	return results
}

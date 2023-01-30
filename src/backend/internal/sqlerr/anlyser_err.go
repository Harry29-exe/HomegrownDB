package sqlerr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/dberr"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"fmt"
)

var AnlsrErr = anlsr{}

type anlsr struct{}

// -------------------------
//      ColumnNotExist
// -------------------------

func (a anlsr) NewColumnNotExist(query node.Query, colName, tableAlias string) ColumnNotExist {
	return ColumnNotExist{
		Query:      query,
		ColName:    colName,
		TableAlias: tableAlias,
	}
}

var _ error = ColumnNotExist{}

type ColumnNotExist struct {
	Query      node.Query
	ColName    string
	TableAlias string
}

func (c ColumnNotExist) Error() string {
	if c.TableAlias == "" {
		return fmt.Sprintf("Could not find column: %s on table: %s",
			c.ColName, c.TableAlias)
	} else {
		return fmt.Sprintf("Could not find column: %s",
			c.ColName)
	}
}

// -------------------------
//      TableNotExist
// -------------------------

func (a anlsr) TableWithAliasNotExist(alias string) TableNotExist {
	return TableNotExist{TableAlias: alias}
}

type TableNotExist struct {
	TableName  string
	TableAlias string
}

func (e TableNotExist) Error() string {
	if e.TableAlias != "" {
		return fmt.Sprintf("table with alias: \"%s\" does not exist", e.TableAlias)
	} else {
		return fmt.Sprintf("table with name: \"%s\" does not exist", e.TableName)
	}
}

// -------------------------
//      IllegalNode
// -------------------------

type IllegalNode struct {
}

// -------------------------
//      TypeMismatch
// -------------------------

func NewTypeMismatch(expectedType rawtype.Tag, actualType rawtype.Tag, value any) TypeMismatch {
	return TypeMismatch{
		ExpectedType: expectedType,
		ActualType:   actualType,
		Value:        value,
	}
}

var _ dberr.DBError = TypeMismatch{}

type TypeMismatch struct {
	ExpectedType rawtype.Tag
	ActualType   rawtype.Tag
	Value        any
}

func (t TypeMismatch) Error() string {
	return fmt.Sprintf("expected type %s but value %+v is of type %s",
		t.ExpectedType.ToStr(), t.Value, t.ActualType.ToStr())
}

func (t TypeMismatch) Area() dberr.Area {
	return dberr.DBSystem
}

func (t TypeMismatch) MsgCanBeReturnedToClient() bool {
	return true
}

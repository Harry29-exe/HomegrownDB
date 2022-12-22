package sqlerr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/dberr"
	"HomegrownDB/dbsystem/hgtype"
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
//      IllegalNode
// -------------------------

type IllegalNode struct {
}

// -------------------------
//      TypeMismatch
// -------------------------

var _ dberr.DBError = TypeMismatch{}

type TypeMismatch struct {
	ExpectedType hgtype.Tag
	ActualType   hgtype.Tag
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

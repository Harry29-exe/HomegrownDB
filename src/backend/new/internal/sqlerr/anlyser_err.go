package sqlerr

import (
	"HomegrownDB/backend/new/internal/node"
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

type IllegalNode struct {
}

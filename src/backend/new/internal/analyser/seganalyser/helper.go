package seganalyser

import (
	"HomegrownDB/backend/new/internal/node"
)

var QueryHelper = queryHelper{}

type queryHelper struct{}

func (queryHelper) findRteByAlias(alias string, query node.Query) node.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Alias != nil && rTable.Alias.AliasName == alias {
			return rTable
		}
	}
	return nil
}

func (queryHelper) findRteWithField(fieldName string, query node.Query) node.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Kind != node.RteRelation {
			continue
		}

		for _, col := range rTable.Ref.Columns() {
			colName := col.Name()
			if colName == fieldName {
				return rTable
			}
		}
	}
	return nil
}

func (queryHelper) findRteWithId(id node.RteID, query node.Query) node.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Id == id {
			return rTable
		}
	}
	return nil
}

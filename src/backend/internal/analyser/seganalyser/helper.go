package seganalyser

import (
	node2 "HomegrownDB/backend/internal/node"
)

var QueryHelper = queryHelper{}

type queryHelper struct{}

func (queryHelper) findRteByAlias(alias string, query node2.Query) node2.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Alias != nil && rTable.Alias.AliasName == alias {
			return rTable
		}
	}
	return nil
}

func (queryHelper) findRteWithField(fieldName string, query node2.Query) node2.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Kind != node2.RteRelation {
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

func (queryHelper) findRteWithId(id node2.RteID, query node2.Query) node2.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Id == id {
			return rTable
		}
	}
	return nil
}

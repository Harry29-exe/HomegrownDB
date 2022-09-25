package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser/pnode"
)

type InsertRows struct{}

func (receiver InsertRows) Analyse(
	rows []pnode.InsertingRow,
	pattern ColumnTypesPattern,
) (anode.InsertRows, error) {
	//todo implement me
	panic("Not implemented")

	//for _, row := range rows {
	//	if !pattern.RowMatches(row) {
	//		//return anode.Values{}, queryerr.ColumnNotExist()
	//	}
	//}

}

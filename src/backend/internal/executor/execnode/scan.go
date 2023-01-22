package execnode

import (
	"HomegrownDB/backend/internal/executor/execnode/exexpr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/tx"
)

type scan struct {
	Plan node.Plan
	Tx   tx.Tx
}

func (s scan) createOutputTuple(internal dpage.Tuple) dpage.Tuple {
	targetList := s.Plan.Plan().TargetList
	exInput := exexpr.ExNodeInput{
		Plan:       s.Plan,
		Internal:   internal,
		LeftInput:  dpage.Tuple{},
		RightInput: dpage.Tuple{},
	}
	patternCols := make([]dpage.ColumnInfo, len(targetList))

	values := make([][]byte, len(targetList))
	for i, targetEntry := range targetList {
		values[i] = exexpr.Execute(targetEntry.ExprToExec, exInput)
		patternCols[i] = dpage.ColumnInfo{
			CType: targetEntry.TypeTag().Type(),
			Name:  targetEntry.ColName,
		}
	}

	return dpage.NewTuple(values, dpage.NewPattern(patternCols), s.Tx)
}

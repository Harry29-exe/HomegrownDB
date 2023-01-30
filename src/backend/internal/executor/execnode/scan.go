package execnode

import (
	"HomegrownDB/backend/internal/executor/execnode/exexpr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

type scan struct {
	Plan node.Plan
	Tx   tx.Tx
}

func (s scan) createOutputTuple(internal page.Tuple) page.Tuple {
	targetList := s.Plan.Plan().TargetList
	exInput := exexpr.ExNodeInput{
		Plan:       s.Plan,
		Internal:   internal,
		LeftInput:  page.Tuple{},
		RightInput: page.Tuple{},
	}
	patternCols := make([]page.PatternCol, len(targetList))

	values := make([][]byte, len(targetList))
	for i, targetEntry := range targetList {
		values[i] = exexpr.Execute(targetEntry.ExprToExec, exInput)
		patternCols[i] = page.PatternCol{
			Type: coltype.NewDefaultColType(targetEntry.TypeTag()),
			Name: targetEntry.ColName,
		}
	}

	return page.NewTuple(values, page.NewPattern(patternCols), s.Tx)
}

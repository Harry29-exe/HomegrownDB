package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var TargetEntry = targetEntry{}

type targetEntry struct{}

func (e targetEntry) Analyse(target pnode.ResultTarget, ctx query.Ctx) (node.TargetEntry, error) {
	val := target.Val

	switch val.Tag() {
	case pnode.TagColumnRef:

	}
}

func (e targetEntry) analyseColumnRef(
	target pnode.ResultTarget,
	colRef pnode.ColumnRef,
	ctx query.Ctx,
) (node.TargetEntry, error) {

}

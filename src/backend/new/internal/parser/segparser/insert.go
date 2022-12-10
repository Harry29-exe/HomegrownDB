package segparser

import (
	"HomegrownDB/backend/new/internal/parser/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Parse(src tkSource, v tkValidator) (pnode.InsertStmt, error) {
	src.Checkpoint()

	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	src.Next()
	insertNode := pnode.NewInsertStmt()

	insertNode.Relation, err = RangeVar.Parse(src, v)
	if err != nil {
		src.Rollback()
		return nil, err
	}

	if err = v.NextIs(token.SpaceBreak); err != nil {
		src.Rollback()
		return insertNode, err
	}

	if err = v.NextIs(token.OpeningParenthesis); err == nil {
		insertNode.Columns, err = ResultTargets.Parse(src, v, ResultTargetInsert)
		if err != nil {
			src.Rollback()
			return insertNode, err
		} else if err = v.CurrentIs(token.SpaceBreak); err != nil {
			src.Rollback()
			return insertNode, err
		}
	} else {
		insertNode.Columns = []pnode.ResultTarget{pnode.NewAStarResultTarget()}
	}

	selectStmt, err := Select.Parse(src, v)
	if err != nil {
		src.Rollback()
		return insertNode, err
	}
	insertNode.SrcNode = selectStmt
	src.CommitAndInitNode(insertNode)
	return insertNode, nil
}

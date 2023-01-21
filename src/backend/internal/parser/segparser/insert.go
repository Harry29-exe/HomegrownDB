package segparser

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	pnode2 "HomegrownDB/backend/internal/pnode"
)

var Insert = insert{}

type insert struct{}

func (i insert) Parse(src tkSource, v tkValidator) (pnode2.InsertStmt, error) {
	src.Checkpoint()

	err := v.CurrentSequence(token.Insert, token.SpaceBreak, token.Into, token.SpaceBreak)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	src.Next()
	insertNode := pnode2.NewInsertStmt()

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
		insertNode.Columns = []pnode2.ResultTarget{pnode2.NewAStarResultTarget()}
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

package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
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

	relation, err := Table.Parse(src, v)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	insertNode.Relation = relation

	err = v.NextSequence(token.SpaceBreak, token.OpeningParenthesis)
	if err == nil {
		resTargets, err := ResultTargets.Parse(src, v)
		if err != nil {
			return insertNode, err
		}
		insertNode.Columns = resTargets
	} else {
		insertNode.Columns = []pnode.ResultTarget{pnode.NewAStarResultTarget()}
	}

	err = v.NextIs(token.SpaceBreak)
	if err != nil {
		src.Rollback()
		return insertNode, err
	}

	src.Next()

	selectStmt, err := Select.Parse(src, v)
	if err != nil {
		return insertNode, err
	}
	insertNode.SrcNode = selectStmt
	return insertNode, nil
}

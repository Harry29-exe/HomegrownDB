package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/backend/new/internal/sqlerr"
)

var Values = values{}

type values struct{}

func (val values) Parse(src tkSource, v tkValidator) ([]pnode.Node, error) {
	nodes := make([]pnode.Node, 0, 10)

	node, err := Value.Parse(src, v)
	if err != nil {
		return nil, err
	}
	nodes = append(nodes, node)

	for val.hasNextVal(v) {
		node, err = Value.Parse(src, v)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (val values) hasNextVal(v tkValidator) bool {
	return v.SkipCurrentSBAnd().CurrentIsAnd(token.Comma).SkipCurrentSB() == nil
}

var Value = value{}

type value struct{}

func (val value) Parse(src tkSource, v tkValidator) (node pnode.Node, err error) {
	src.Checkpoint()

	if val.isFunction(v) {
		src.Rollback()
		//todo implement me
		panic("Not implemented")

	} else if val.isColumnRef(v) {
		src.Rollback()
		node, err = val.parseColumnRef(src, v)

	} else {
		node, err = val.parseConst(src, v)
	}

	if err != nil {
		src.Rollback()
	} else {
		src.Commit()
	}
	return
}

func (val value) isFunction(v tkValidator) bool {
	return v.CurrentSequence(token.Identifier, token.OpeningParenthesis) == nil
}

func (val value) isColumnRef(v tkValidator) bool {
	return v.CurrentSequence(token.Identifier, token.SpaceBreak) == nil ||
		v.CurrentSequence(token.Identifier, token.Dot, token.Identifier) == nil
}

func (val value) parseColumnRef(src tkSource, v tkValidator) (pnode.ColumnRef, error) {
	identifier1 := src.Current()
	if tk := src.Next(); tk.Code() == token.SpaceBreak {
		return pnode.NewColumnRef(identifier1.Value(), ""), nil
	} else {
		cRef := pnode.NewColumnRef(src.Next().Value(), identifier1.Value())
		src.Next()
		return cRef, nil
	}
}

func (val value) parseConst(src tkSource, v tkValidator) (pnode.AConst, error) {
	src.Checkpoint()
	tk := src.Current()
	var aConst pnode.AConst
	var err error

	switch tk.Code() {
	case token.SqlTextValue:
		textTk := tk.(*token.SqlTextValueToken)
		aConst, err = pnode.NewAConstStr(textTk.InnerStr), nil
	case token.Float:
		floatTk := tk.(*token.FloatToken)
		aConst, err = pnode.NewAConstFloat(floatTk.Float), nil
	case token.Integer:
		intTk := tk.(*token.IntegerToken)
		aConst, err = pnode.NewAConstInt(intTk.Int), nil
	default:
		src.Rollback()
		return nil, sqlerr.NewSyntaxError(
			"Const value str/int/float",
			token.ToString(tk.Code()),
			src,
		)
	}

	if err != nil {
		return nil, err
	}
	src.Next()
	src.CommitAndInitNode(aConst)
	return aConst, nil
}

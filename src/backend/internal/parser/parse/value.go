package parse

import (
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
)

var Values = values{}

type values struct{}

func (val values) Parse(src tkSource, v tkValidator) ([]pnode2.Node, error) {
	nodes := make([]pnode2.Node, 0, 10)

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
	return v.SkipCurrentSBAnd().CurrentIsAnd(token2.Comma).SkipCurrentSB() == nil
}

var Value = value{}

type value struct{}

func (val value) Parse(src tkSource, v tkValidator) (node pnode2.Node, err error) {
	src.Checkpoint()

	if val.isFunction(v) {
		src.Rollback()
		//todo implement me
		panic("Not implemented")

	} else if val.isColumnRef(v) {
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
	return v.CurrentSequence(token2.Identifier, token2.OpeningParenthesis) == nil
}

func (val value) isColumnRef(v tkValidator) bool {
	return v.CurrentSequence(token2.Identifier) == nil
}

func (val value) parseColumnRef(src tkSource, v tkValidator) (pnode2.ColumnRef, error) {
	identifier1 := src.Current()
	if tk := src.Next(); tk.Code() == token2.Dot {
		cRef := pnode2.NewColumnRef(src.Next().Value(), identifier1.Value())
		src.Next()
		return cRef, nil
	} else {
		return pnode2.NewColumnRef(identifier1.Value(), ""), nil
	}
}

func (val value) parseConst(src tkSource, v tkValidator) (pnode2.AConst, error) {
	src.Checkpoint()
	tk := src.Current()
	var aConst pnode2.AConst
	var err error

	switch tk.Code() {
	case token2.SqlTextValue:
		textTk := tk.(*token2.SqlTextValueToken)
		aConst, err = pnode2.NewAConstStr(textTk.InnerStr), nil
	case token2.Float:
		floatTk := tk.(*token2.FloatToken)
		aConst, err = pnode2.NewAConstFloat(floatTk.Float), nil
	case token2.Integer:
		intTk := tk.(*token2.IntegerToken)
		aConst, err = pnode2.NewAConstInt(intTk.Int), nil
	default:
		src.Rollback()
		return nil, sqlerr.NewSyntaxError(
			"Const value str/int/float",
			token2.ToString(tk.Code()),
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

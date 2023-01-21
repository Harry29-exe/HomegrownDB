package typanlr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype"
)

type FutureType struct {
	TypeTag  hgtype.Tag
	TypeArgs hgtype.Args
}

func (f *FutureType) UpdateType(expr node2.Expr) error {
	if f.TypeTag != expr.Type() {
		return sqlerr.TypeMismatch{
			ExpectedType: f.TypeTag,
			ActualType:   expr.Type(),
			Value:        expr,
		}
	}
	switch expr.Tag() {
	case node2.TagConst:
		return f.updateByConst(expr.(node2.Const))
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (f *FutureType) updateByConst(expr node2.Const) error {
	switch f.TypeTag {
	case hgtype.TypeInt8:
		return nil
	case hgtype.TypeStr:
		return f.updateStr(expr.Val)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (f *FutureType) updateStr(str []byte) error {
	strLen := hgtype.StrUtils.StrLen(str)
	if f.TypeArgs.Length < strLen {
		f.TypeArgs.Length = strLen - 4 //string len - 4 bytes of header
	}
	if !f.TypeArgs.UTF8 && !hgtype.StrUtils.IsASCII(str) {
		f.TypeArgs.UTF8 = true
	}

	return nil
}

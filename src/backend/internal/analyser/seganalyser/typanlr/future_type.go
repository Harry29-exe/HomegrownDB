package typanlr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/sqlerr"
	"HomegrownDB/dbsystem/hgtype/rawtype"
)

type FutureType struct {
	TypeTag  rawtype.Tag
	TypeArgs rawtype.Args
}

func (f *FutureType) UpdateType(expr node2.Expr) error {
	if f.TypeTag != expr.TypeTag() {
		return sqlerr.TypeMismatch{
			ExpectedType: f.TypeTag,
			ActualType:   expr.TypeTag(),
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
	case rawtype.TypeInt8:
		return nil
	case rawtype.TypeStr:
		return f.updateStr(expr.Val)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (f *FutureType) updateStr(str []byte) error {
	strLen := rawtype.StrUtils.StrLen(str)
	if f.TypeArgs.Length < strLen {
		f.TypeArgs.Length = strLen - 4 //string len - 4 bytes of header
	}
	if !f.TypeArgs.UTF8 && !rawtype.StrUtils.IsASCII(str) {
		f.TypeArgs.UTF8 = true
	}

	return nil
}

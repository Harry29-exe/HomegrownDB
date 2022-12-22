package typa

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/hgtype"
)

func CreateFutureType(expr node.Expr) FutureType {
	switch expr.Tag() {
	case node.TagConst:
		return createFTFromConst(expr.(node.Const))
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func createFTFromConst(expr node.Const) FutureType {
	switch expr.Type() {
	case hgtype.TypeInt8:
		return FutureType{
			TypeTag:  hgtype.TypeInt8,
			TypeArgs: hgtype.Args{},
		}
	case hgtype.TypeStr:
		args := hgtype.Args{}
		args.Length = uint32(len(expr.Val))
		args.UTF8 = !hgtype.StrUtils.IsASCII(expr.Val)
		args.VarLen = true

		return FutureType{
			TypeTag:  hgtype.TypeStr,
			TypeArgs: hgtype.Args{},
		}
	default:
		//todo implement me
		panic("Not implemented")
	}
}

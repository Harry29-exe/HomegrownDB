package typanlr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hgtype/rawtype"
)

func CreateFutureType(expr node2.Expr) FutureType {
	switch expr.Tag() {
	case node2.TagConst:
		return createFTFromConst(expr.(node2.Const))
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func createFTFromConst(expr node2.Const) FutureType {
	switch expr.TypeTag() {
	case rawtype.TypeInt8:
		return FutureType{
			TypeTag:  rawtype.TypeInt8,
			TypeArgs: rawtype.Args{},
		}
	case rawtype.TypeStr:
		args := rawtype.Args{
			Length: uint32(len(expr.Val)),
			UTF8:   !rawtype.StrUtils.IsASCII(expr.Val),
			VarLen: true,
		}

		return FutureType{
			TypeTag:  rawtype.TypeStr,
			TypeArgs: args,
		}
	default:
		//todo implement me
		panic("Not implemented")
	}
}

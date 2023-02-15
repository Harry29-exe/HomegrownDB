package typanlr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hgtype/rawtype"
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
	switch expr.TypeTag() {
	case rawtype.TypeInt8:
		return FutureType{
			TypeTag: rawtype.TypeInt8,
			TypeArgs: rawtype.Args{
				Length:   8,
				VarLen:   false,
				Nullable: true,
			},
		}
	case rawtype.TypeStr:
		args := rawtype.Args{
			Length:   len(expr.Val),
			UTF8:     !rawtype.StrUtils.IsASCII(expr.Val),
			VarLen:   true,
			Nullable: true,
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

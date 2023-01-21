package typanlr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hgtype"
)

func CreateFutureTypes(exprs []node.Expr) FutureTypes {
	types := make([]FutureType, len(exprs))
	for i, expr := range exprs {
		types[i] = CreateFutureType(expr)
	}
	return FutureTypes{
		Types: types,
	}
}

type FutureTypes struct {
	Types []FutureType
}

func (f *FutureTypes) UpdateTypes(exprs []node.Expr) error {
	if len(exprs) != len(f.Types) {
		panic("illegal state: length of f.Types does not matches length of exprs")
	}

	for i, expr := range exprs {
		err := f.Types[i].UpdateType(expr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FutureTypes) CreateTypes() []hgtype.TypeData {
	types := make([]hgtype.TypeData, len(f.Types))
	for i, ft := range f.Types {
		types[i] = hgtype.NewTypeData(ft.TypeTag, ft.TypeArgs)
	}
	return types
}

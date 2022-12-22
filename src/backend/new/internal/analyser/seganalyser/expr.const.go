package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/dbsystem/hgtype"
)

func (ex exprAnalyser) AnalyseConst(aConst pnode.AConst, query node.Query, ctx anlsr.Ctx) (node.Const, error) {
	switch aConst.Type {
	case pnode.AConstInt:
		return node.NewConstInt8(aConst.Int, hgtype.Args{}), nil
	case pnode.AConstStr:
		return node.NewConstStr(aConst.Str, hgtype.Args{Length: uint32(len(aConst.Str))})
	case pnode.AConstFloat:
		//todo implement me
		panic("Not implemented")
		//return node.NewConst(hgtype.TypeFloat8, aConst.Float), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}

//
//type Consts struct {
//	Values  [][]node.Const
//	Pattern []hgtype.HGType
//}
//
//func (ex exprAnalyser) AnalyseConsts(
//	aConst [][]pnode.AConst,
//	query node.Query,
//	ctx anlsr.Ctx,
//) (
//	Consts,
//	error,
//) {
//	if len(aConst) < 1 {
//		return Consts{}, errors.New("values array is empty") //todo better err
//	}
//	rowLen := len(aConst[0])
//	types := make([]hgtype.HGType, rowLen)
//
//	firstRow := aConst[0]
//	for col := 0; col < len(aConst); col++ {
//		constNode, err := ex.AnalyseConst(firstRow[col], query, ctx)
//
//	}
//	for row := 0; row < len(aConst); row++ {
//
//	}
//}
//
//var ConstAnalyser = constsAnalyser{}
//
//type constsAnalyser struct{}
//
//func (c constsAnalyser) analyseFirstRow(firstRow []pnode.AConst, rowCount int) (Consts, error) {
//	consts := Consts{
//		Values:  make([][]node.Const, rowCount),
//		Pattern: make([]hgtype.HGType, len(firstRow)),
//	}
//	colCount := len(firstRow)
//	for row := 0; row < rowCount; row++ {
//		consts.Values[row] = make([]node.Const, colCount)
//	}
//
//	for col := 0; col < len(firstRow); col++ {
//		constNode, err := c.analyseConst(firstRow[col])
//		if err != nil {
//			return Consts{}, err
//		}
//		consts.Values[0][col] = constNode
//		consts.Pattern[col] = constNode.TypeTag
//	}
//	return consts, nil
//}
//
//func (c constsAnalyser) analyseRows(rows [][]pnode.AConst, consts Consts) error {
//	for row := 0; row < len(rows); row++ {
//		for col := 0; col < len(consts.Pattern); col++ {
//
//		}
//	}
//}
//
//func (c constsAnalyser) analyseRowColumn(aConst pnode.AConst, hgType hgtype.HGType) (node.Const, error) {
//	switch aConst.TypeTag {
//	case pnode.AConstInt:
//		return c.convertInt(aConst, hgType)
//	case pnode.AConstStr:
//
//	}
//}
//
//func (c constsAnalyser) analyseConst(aConst pnode.AConst) (node.Const, error) {
//	switch aConst.TypeTag {
//	case pnode.AConstInt:
//		colType := hgtype.NewInt8(hgtype.Int8Args{})
//		return node.NewConst(colType, hgtype.Int8Serialize(aConst.Int)), nil
//	case pnode.AConstStr:
//		colType := hgtype.NewStr(hgtype.StrArgs{Length: uint32(len(aConst.Str))})
//		return node.NewConst(colType, hgtype.StrSerializeInput(aConst.Str)), nil
//	case pnode.AConstFloat:
//		//todo implement me
//		panic("Not implemented")
//		//return node.NewConst(hgtype.TypeFloat8, aConst.Float), nil
//	default:
//		//todo implement me
//		panic("Not implemented")
//	}
//}
//
//func (c constsAnalyser) convertInt(aConst pnode.AConst, constType hgtype.HGType) (node.Const, error) {
//	if constType.TypeTag == hgtype.TypeInt8 {
//		node.NewConst(constType, hgtype.Int8Serialize(aConst.Int))
//	}
//	return nil, fmt.Errorf("can not convert int into %s", constType.TypeTag.ToStr()) //todo better err
//}
//
//func (c constsAnalyser) convertStr(aConst pnode.AConst, constType hgtype.HGType) (node.Const, error) {
//	switch constType.TypeTag {
//	case hgtype.TypeStr:
//		if constType. {
//
//		}
//	}
//
//	return nil, c.canNotConvertErr(aConst, constType)
//}
//
//func (c constsAnalyser) canNotConvertErr(fromNode pnode.AConst, constType hgtype.HGType) error {
//	return fmt.Errorf("can not convert %+v (type: %s) into %s",
//		fromNode,
//		fromNode.TypeTag.ToStr(),
//		constType.TypeTag.ToStr(),
//	) //todo better err
//}

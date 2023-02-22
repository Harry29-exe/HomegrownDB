package internal

import (
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/lib/bparse"
)

func GetString(order reldef.Order, tuple page.RTuple) string {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return string(rawtype.TypeOp.GetData(value, args))
}

func GetInt8(order reldef.Order, tuple page.RTuple) int64 {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Int8(rawtype.TypeOp.GetData(value, args))
}

func GetBool(order reldef.Order, tuple page.RTuple) bool {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Bool(rawtype.TypeOp.GetData(value, args))
}

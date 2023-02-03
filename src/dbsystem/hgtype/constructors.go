package hgtype

import "HomegrownDB/dbsystem/hgtype/rawtype"

func NewStr(args Args) ColumnType {
	return ColumnType{
		ColType: rawtype.Str{},
		ColTag:  rawtype.TypeStr,
		ColArgs: args,
	}
}

func NewInt8(args Args) ColumnType {
	args.Length = 8
	return ColumnType{
		ColType: rawtype.Int8{},
		ColArgs: args,
	}
}

func NewBool(args Args) ColumnType {
	args.Length = 1
	return ColumnType{
		ColType: rawtype.Bool{},
		ColTag:  rawtype.TypeBool,
		ColArgs: args,
	}
}

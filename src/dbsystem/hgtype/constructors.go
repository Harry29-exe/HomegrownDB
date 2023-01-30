package hgtype

import "HomegrownDB/dbsystem/hgtype/rawtype"

func NewStr(args Args) ColumnType {
	return ColumnType{
		Type: rawtype.Str{},
		Tag:  rawtype.TypeStr,
		Args: args,
	}
}

func NewInt8(args Args) ColumnType {
	return ColumnType{
		Type: rawtype.Int8{},
		Args: args,
	}
}

func NewBool(args Args) ColumnType {
	return ColumnType{
		Type: rawtype.Bool{},
		Tag:  rawtype.TypeBool,
		Args: args,
	}
}

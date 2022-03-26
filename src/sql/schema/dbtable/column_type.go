package dbtable

// ColumnType describes column data structure as int or string
type ColumnType struct {
	Code ColumnTypeCode
	// LenPrefixSize describes how long is column prefix describing
	// how long column data is, if LenPrefixSize is 0 this mean column is
	// fixed size and does not have len prefix
	LenPrefixSize uint8
	// IsFixedSize true then this describes max size of column
	// where 0 means no max size otherwise it describes size od column
	ByteLen uint32

	LobStatus IsLOB
}

// IsLOB enum to indicate if given column type can be saved as large object
type IsLOB = uint8

const (
	NEVER IsLOB = iota
	SOMETIMES
	ALWAYS
)

func GetColumnType(code ColumnTypeCode, args []int32) *ColumnType {
	factoryFun, ok := columnTypeMap[code]
	if !ok {
		panic("There is not column type with code: " + code)
	}

	return factoryFun(args)
}

type ColumnTypeCode = string

const (
	Int2    ColumnTypeCode = "int2"
	Int4    ColumnTypeCode = "int4"
	Int8    ColumnTypeCode = "int8"
	Char    ColumnTypeCode = "char"
	Varchar ColumnTypeCode = "varchar"
	Bool    ColumnTypeCode = "bool"
	BLOB    ColumnTypeCode = "blob"
)

type columnTypeFactory = func(args []int32) *ColumnType

var columnTypeMap = map[ColumnTypeCode]columnTypeFactory{
	Int2: func(args []int32) *ColumnType {
		return createSimpleColumnType(Int2, 2)
	},
	Int4: func(args []int32) *ColumnType {
		return createSimpleColumnType(Int4, 4)
	},
	Int8: func(args []int32) *ColumnType {
		return createSimpleColumnType(Int8, 8)
	},
	Char: func(args []int32) *ColumnType {
		var strLen uint32 = 1
		if len(args) > 0 {
			strLen = uint32(args[0])
		}

		return &ColumnType{
			Code:          Char,
			LenPrefixSize: 0,
			ByteLen:       strLen,
			LobStatus:     SOMETIMES,
		}
	},
	Varchar: func(args []int32) *ColumnType {
		var strLen uint32 = 1
		if len(args) > 1 {
			strLen = uint32(args[0])
		}

		return &ColumnType{
			Code:          Varchar,
			LenPrefixSize: 0,
			ByteLen:       strLen,
			LobStatus:     SOMETIMES,
		}
	},
	Bool: func(args []int32) *ColumnType {
		return createSimpleColumnType(Bool, 1)
	},
	BLOB: func(args []int32) *ColumnType {
		return &ColumnType{
			Code:          BLOB,
			LenPrefixSize: 0,
			ByteLen:       0,
			LobStatus:     ALWAYS,
		}
	},
}

func createSimpleColumnType(code ColumnTypeCode, byteLen uint32) *ColumnType {
	return &ColumnType{
		Code:          code,
		LenPrefixSize: 0,
		ByteLen:       byteLen,
		LobStatus:     NEVER,
	}
}

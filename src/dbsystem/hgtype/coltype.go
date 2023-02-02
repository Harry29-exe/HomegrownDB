package hgtype

import (
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hgtype/rawtype"
)

var (
	_ CTOperations = ColumnType{}
	_ CTReader     = ColumnType{}
	_ CTWriter     = ColumnType{}
	_ CTDebug      = ColumnType{}
)

func NewColType(tag rawtype.Tag, args rawtype.Args) ColumnType {
	t := tag.Type()

	return ColumnType{
		ColType: t,
		ColTag:  tag,
		ColArgs: args,
	}
}

func NewDefaultColType(tag rawtype.Tag) ColumnType {
	var t rawtype.Type
	args := rawtype.Args{
		Nullable: true,
	}

	switch tag {
	case rawtype.TypeStr:
		t = rawtype.Str{}
		args.UTF8 = true
		args.Length = uint32(rawtype.UnknownVarLenSize)
		args.VarLen = true
	case rawtype.TypeInt8:
		t = rawtype.Int8{}
	default:
		//todo implement me
		panic("Not implemented")
	}

	return ColumnType{
		ColType: t,
		ColTag:  tag,
		ColArgs: args,
	}
}

type ColumnType struct {
	ColType rawtype.Type
	ColTag  rawtype.Tag
	ColArgs rawtype.Args
}

func (t ColumnType) Type() rawtype.Type {
	return t.ColType
}
func (t ColumnType) Tag() rawtype.Tag {
	return t.ColTag
}

func (t ColumnType) Args() Args {
	return t.ColArgs
}

func (t ColumnType) Validate(value rawtype.Value) rawtype.ValidateResult {
	return t.ColType.Validate(t.ColArgs, value)
}

// -------------------------
//      CTReader
// -------------------------

func (t ColumnType) Skip(data []byte) []byte {
	return t.ColType.Skip(data)
}

func (t ColumnType) Copy(dest []byte, data []byte) (copiedBytes int) {
	return t.ColType.Copy(dest, data)
}

func (t ColumnType) IsToastPtr(data []byte) bool {
	return t.ColType.IsToastPtr(data)
}

func (t ColumnType) Value(data []byte) (value []byte) {
	return t.ColType.Value(data)
}

func (t ColumnType) ValueAndSkip(data []byte) (value, next []byte) {
	return t.ColType.ValueAndSkip(data)
}

// -------------------------
//      CTWriter
// -------------------------

func (t ColumnType) WriteValue(writer rawtype.UniWriter, value rawtype.Value) error {
	return t.ColType.WriteValue(writer, value, t.ColArgs)
}

// -------------------------
//      CTOperations
// -------------------------

func (t ColumnType) Equal(v1, v2 []byte) bool {
	return t.ColType.Equal(v1, v2)
}

func (t ColumnType) Cmp(v1, v2 []byte) int {
	return t.ColType.Cmp(v1, v2)
}

// -------------------------
//      CTDebug
// -------------------------

func (t ColumnType) ToStr(val []byte) string {
	return t.ColType.ToStr(val)
}

func (t ColumnType) Rand(r random.Random) []byte {
	return t.ColType.Rand(t.ColArgs, r)
}

package hgtype

import (
	"HomegrownDB/common/bparse"
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
		Type: t,
		Tag:  tag,
		Args: args,
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
		Type: t,
		Tag:  tag,
		Args: args,
	}
}

func NewStr(args rawtype.Args) ColumnType {
	return ColumnType{
		Type: rawtype.Str{},
		Tag:  rawtype.TypeStr,
		Args: args,
	}
}

func NewInt8(args rawtype.Args) ColumnType {
	return ColumnType{
		Type: rawtype.Int8{},
		Args: args,
	}
}

type ColumnType struct {
	Type rawtype.Type
	Tag  rawtype.Tag
	Args rawtype.Args
}

func SerializeTypeData(typeData ColumnType, s *bparse.Serializer) {
	s.Uint8(uint8(typeData.Type.Tag()))
	rawtype.SerializeArgs(typeData.Args, s)
}

func DeserializeTypeData(d *bparse.Deserializer) ColumnType {
	tag := rawtype.Tag(d.Uint8())
	args := rawtype.DeserializeArgs(d)
	return NewColType(tag, args)
}

func (t ColumnType) Validate(value rawtype.Value) rawtype.ValidateResult {
	return t.Type.Validate(t.Args, value)
}

// -------------------------
//      CTReader
// -------------------------

func (t ColumnType) Skip(data []byte) []byte {
	return t.Type.Skip(data)
}

func (t ColumnType) Copy(dest []byte, data []byte) (copiedBytes int) {
	return t.Type.Copy(dest, data)
}

func (t ColumnType) IsToastPtr(data []byte) bool {
	return t.Type.IsToastPtr(data)
}

func (t ColumnType) Value(data []byte) (value []byte) {
	return t.Type.Value(data)
}

func (t ColumnType) ValueAndSkip(data []byte) (value, next []byte) {
	return t.Type.ValueAndSkip(data)
}

// -------------------------
//      CTWriter
// -------------------------

func (t ColumnType) WriteValue(writer rawtype.UniWriter, value rawtype.Value) error {
	return t.Type.WriteValue(writer, value, t.Args)
}

// -------------------------
//      CTOperations
// -------------------------

func (t ColumnType) Equal(v1, v2 []byte) bool {
	return t.Type.Equal(v1, v2)
}

func (t ColumnType) Cmp(v1, v2 []byte) int {
	return t.Type.Cmp(v1, v2)
}

// -------------------------
//      CTDebug
// -------------------------

func (t ColumnType) ToStr(val []byte) string {
	return t.Type.ToStr(val)
}

func (t ColumnType) Rand(r random.Random) []byte {
	return t.Type.Rand(t.Args, r)
}

package coltype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
	"HomegrownDB/dbsystem/hgtype"
)

var (
	_ Operations = ColumnType{}
	_ Reader     = ColumnType{}
	_ Writer     = ColumnType{}
	_ Debug      = ColumnType{}
)

func NewColType(tag hgtype.Tag, args hgtype.Args) ColumnType {
	t := tag.Type()

	return ColumnType{
		Type: t,
		Tag:  tag,
		Args: args,
	}
}

func NewDefaultColType(tag hgtype.Tag) ColumnType {
	var t hgtype.Type
	args := hgtype.Args{
		Nullable: true,
	}

	switch tag {
	case hgtype.TypeStr:
		t = hgtype.Str{}
		args.UTF8 = true
		args.Length = uint32(hgtype.UnknownVarLenSize)
		args.VarLen = true
	case hgtype.TypeInt8:
		t = hgtype.Int8{}
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

func NewStr(args hgtype.Args) ColumnType {
	return ColumnType{
		Type: hgtype.Str{},
		Tag:  hgtype.TypeStr,
		Args: args,
	}
}

func NewInt8(args hgtype.Args) ColumnType {
	return ColumnType{
		Type: hgtype.Int8{},
		Args: args,
	}
}

type ColumnType struct {
	Type hgtype.Type
	Tag  hgtype.Tag
	Args hgtype.Args
}

func SerializeTypeData(typeData ColumnType, s *bparse.Serializer) {
	s.Uint8(uint8(typeData.Type.Tag()))
	hgtype.SerializeArgs(typeData.Args, s)
}

func DeserializeTypeData(d *bparse.Deserializer) ColumnType {
	tag := hgtype.Tag(d.Uint8())
	args := hgtype.DeserializeArgs(d)
	return NewColType(tag, args)
}

func (t ColumnType) Validate(value hgtype.Value) hgtype.ValidateResult {
	return t.Type.Validate(t.Args, value)
}

// -------------------------
//      Reader
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
//      Writer
// -------------------------

func (t ColumnType) WriteValue(writer hgtype.UniWriter, value hgtype.Value) error {
	return t.Type.WriteValue(writer, value, t.Args)
}

// -------------------------
//      Operations
// -------------------------

func (t ColumnType) Equal(v1, v2 []byte) bool {
	return t.Type.Equal(v1, v2)
}

func (t ColumnType) Cmp(v1, v2 []byte) int {
	return t.Type.Cmp(v1, v2)
}

// -------------------------
//      Debug
// -------------------------

func (t ColumnType) ToStr(val []byte) string {
	return t.Type.ToStr(val)
}

func (t ColumnType) Rand(r random.Random) []byte {
	return t.Type.Rand(t.Args, r)
}

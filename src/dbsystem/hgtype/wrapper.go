package hgtype

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/random"
)

var (
	_ WrapperOp  = TypeData{}
	_ Operations = TypeData{}
	_ Reader     = TypeData{}
	_ Writer     = TypeData{}
	_ Debug      = TypeData{}
)

func NewTypeData(tag Tag, args Args) TypeData {
	t := tag.Type()

	return TypeData{
		Type: t,
		Tag:  tag,
		Args: args,
	}
}

func NewTypeDataWithDefaultArgs(tag Tag) TypeData {
	var t Type
	args := Args{
		Nullable: true,
	}

	switch tag {
	case TypeStr:
		t = Str{}
		args.UTF8 = true
		args.Length = uint32(UnknownVarLenSize)
		args.VarLen = true
	case TypeInt8:
		t = Int8{}
	default:
		//todo implement me
		panic("Not implemented")
	}

	return TypeData{
		Type: t,
		Tag:  tag,
		Args: args,
	}
}

func NewStr(args Args) TypeData {
	return TypeData{
		Type: Str{},
		Tag:  TypeStr,
		Args: args,
	}
}

func NewInt8(args Args) TypeData {
	return TypeData{
		Type: Int8{},
		Tag:  TypeInt8,
		Args: args,
	}
}

type TypeData struct {
	Type Type
	Tag  Tag
	Args Args
}

func (w TypeData) Validate(value []byte) error {
	return w.Type.Validate(w.Args, value)
}

func SerializeTypeData(typeData TypeData, s *bparse.Serializer) {
	s.Uint8(uint8(typeData.Tag))
	SerializeArgs(typeData.Args, s)
}

func DeserializeTypeData(d *bparse.Deserializer) TypeData {
	tag := Tag(d.Uint8())
	args := DeserializeArgs(d)
	return NewTypeData(tag, args)
}

// -------------------------
//      WrapperOp
// -------------------------

func (w TypeData) TypeEqual(wrapper TypeData) bool {
	return w.Tag == wrapper.Tag && w.Args == wrapper.Args
}

// -------------------------
//      Reader
// -------------------------

func (w TypeData) Skip(data []byte) []byte {
	return w.Type.Skip(data)
}

func (w TypeData) Copy(dest []byte, data []byte) (copiedBytes int) {
	return w.Type.Copy(dest, data)
}

func (w TypeData) IsToastPtr(data []byte) bool {
	return w.Type.IsToastPtr(data)
}

func (w TypeData) Value(data []byte) (value []byte) {
	return w.Type.Value(data)
}

func (w TypeData) ValueAndSkip(data []byte) (value, next []byte) {
	return w.Type.ValueAndSkip(data)
}

// -------------------------
//      Writer
// -------------------------

func (w TypeData) WriteTuple(dest []byte, value []byte) int {
	return w.Type.WriteTuple(dest, value)
}

// -------------------------
//      Operations
// -------------------------

func (w TypeData) Equal(v1, v2 []byte) bool {
	return w.Type.Equal(v1, v2)
}

func (w TypeData) Cmp(v1, v2 []byte) int {
	return w.Type.Cmp(v1, v2)
}

// -------------------------
//      Debug
// -------------------------

func (w TypeData) ToStr(val []byte) string {
	return w.Type.ToStr(val)
}

func (w TypeData) Rand(r random.Random) []byte {
	return w.Type.Rand(w.Args, r)
}

package hgtype

import (
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
	var t Type
	switch tag {
	case TypeStr:
		t = Str{}
	case TypeInt8:
		t = Int8{}
	default:
		//todo implement me
		panic("Not implemented")
	}

	return TypeData{
		t:    t,
		Tag:  tag,
		Args: args,
	}
}

func NewStr(args Args) TypeData {
	return TypeData{
		t:    Str{},
		Tag:  TypeStr,
		Args: args,
	}
}

func NewInt8(args Args) TypeData {
	return TypeData{
		t:    Int8{},
		Tag:  TypeInt8,
		Args: args,
	}
}

type TypeData struct {
	t    Type
	Tag  Tag
	Args Args
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
	return w.t.Skip(w.Args, data)
}

func (w TypeData) Copy(dest []byte, data []byte) (copiedBytes int) {
	return w.t.Copy(w.Args, dest, data)
}

func (w TypeData) IsToastPtr(data []byte) bool {
	return w.t.IsToastPtr(w.Args, data)
}

func (w TypeData) Value(data []byte) (value []byte) {
	return w.t.Value(w.Args, data)
}

func (w TypeData) ValueAndSkip(data []byte) (value, next []byte) {
	return w.t.ValueAndSkip(w.Args, data)
}

// -------------------------
//      Writer
// -------------------------

func (w TypeData) WriteTuple(dest []byte, value []byte) int {
	return w.t.WriteTuple(w.Args, dest, value)
}

// -------------------------
//      Operations
// -------------------------

func (w TypeData) Equal(v1, v2 []byte) bool {
	return w.t.Equal(w.Args, v1, v2)
}

func (w TypeData) Cmp(v1, v2 []byte) int {
	return w.t.Cmp(w.Args, v1, v2)
}

// -------------------------
//      Debug
// -------------------------

func (w TypeData) ToStr(val []byte) string {
	return w.t.ToStr(w.Args, val)
}

func (w TypeData) Rand(r random.Random) []byte {
	return w.t.Rand(w.Args, r)
}

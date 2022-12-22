package hgtype

import (
	"HomegrownDB/common/random"
)

var (
	_ WrapperOp  = Wrapper{}
	_ Operations = Wrapper{}
	_ Reader     = Wrapper{}
	_ Writer     = Wrapper{}
	_ Debug      = Wrapper{}
)

func NewWrapper(tag TypeTag, args Args) Wrapper {
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

	return Wrapper{
		t:    t,
		Tag:  tag,
		Args: args,
	}
}

func NewStr(args Args) Wrapper {
	return Wrapper{
		t:    Str{},
		Tag:  TypeStr,
		Args: args,
	}
}

func NewInt8(args Args) Wrapper {
	return Wrapper{
		t:    Int8{},
		Tag:  TypeInt8,
		Args: args,
	}
}

type Wrapper struct {
	t    Type
	Tag  TypeTag
	Args Args
}

// -------------------------
//      WrapperOp
// -------------------------

func (w Wrapper) TypeEqual(wrapper Wrapper) bool {
	return w.Tag == wrapper.Tag && w.Args == wrapper.Args
}

// -------------------------
//      Reader
// -------------------------

func (w Wrapper) Skip(data []byte) []byte {
	return w.t.Skip(w.Args, data)
}

func (w Wrapper) Copy(dest []byte, data []byte) (copiedBytes int) {
	return w.t.Copy(w.Args, dest, data)
}

func (w Wrapper) IsToastPtr(data []byte) bool {
	return w.t.IsToastPtr(w.Args, data)
}

func (w Wrapper) Value(data []byte) (value []byte) {
	return w.t.Value(w.Args, data)
}

func (w Wrapper) ValueAndSkip(data []byte) (value, next []byte) {
	return w.t.ValueAndSkip(w.Args, data)
}

// -------------------------
//      Writer
// -------------------------

func (w Wrapper) WriteTuple(dest []byte, value []byte) int {
	return w.t.WriteTuple(w.Args, dest, value)
}

// -------------------------
//      Operations
// -------------------------

func (w Wrapper) Equal(v1, v2 []byte) bool {
	return w.t.Equal(w.Args, v1, v2)
}

func (w Wrapper) Cmp(v1, v2 []byte) int {
	return w.t.Cmp(w.Args, v1, v2)
}

// -------------------------
//      Debug
// -------------------------

func (w Wrapper) ToStr(val []byte) string {
	return w.t.ToStr(w.Args, val)
}

func (w Wrapper) Rand(r random.Random) []byte {
	return w.t.Rand(w.Args, r)
}

package pnode

// -------------------------
//      ColumnDef
// -------------------------

type ColumnDef = *columnDef

func NewColumnDef(name string, typeName string, args []TypeArg) ColumnDef {
	return &columnDef{
		Name: name,
		Type: typeName,
		Args: args,
	}
}

var _ Node = &columnDef{}

type columnDef struct {
	node
	Name string
	Type string
	Args []TypeArg
}

func (c ColumnDef) Equal(node Node) bool {
	if nodesEqNil(c, node) {
		return true
	} else if !basicNodeEqual(c, node) {
		return false
	}
	raw := node.(ColumnDef)
	return c.Name == raw.Name &&
		c.Type == raw.Type &&
		cmpNodeArray(c.Args, raw.Args)
}

// -------------------------
//      TypeArg
// -------------------------

func NewArgLength(len int) TypeArg {
	return &typeArg{
		node: node{
			tag: TagTypeArg,
		},
		Arg: ArgTypeLength,
		Val: len,
	}
}

type TypeArg = *typeArg

var _ Node = &typeArg{}

type typeArg struct {
	node
	Arg ArgType
	Val any
}

func (t TypeArg) Equal(node Node) bool {
	if nodesEqNil(t, node) {
		return true
	} else if !basicNodeEqual(t, node) {
		return false
	}
	raw := node.(TypeArg)
	return t.Arg == raw.Arg &&
		t.Val == raw.Val
}

type ArgType uint8

const (
	ArgTypeLength ArgType = iota
)

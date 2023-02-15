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
	return &typeArgValue{
		node: node{
			tag: TagTypeArg,
		},
		Arg: TypeArgTypeLength,
		Val: len,
	}
}

type TypeArg = *typeArgValue

var _ Node = &typeArgValue{}

type typeArgValue struct {
	node
	Arg TypeArgType
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

type TypeArgType uint8

const (
	TypeArgTypeLength TypeArgType = iota
	TypeArgTypeNullable
)

func (t TypeArgType) ToString() string {
	return []string{
		"TypeArgTypeLength",
		"TypeArgTypeNullable",
	}[t]
}

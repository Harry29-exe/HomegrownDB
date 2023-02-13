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

type columnDef struct {
	node
	Name string
	Type string
	Args []TypeArg
}

// -------------------------
//      TypeArg
// -------------------------

func NewArgLength(len int) TypeArg {
	return typeArg{
		Arg: ArgTypeLength,
		Val: len,
	}
}

type TypeArg = typeArg

type typeArg struct {
	Arg ArgType
	Val any
}

type ArgType uint8

const (
	ArgTypeLength ArgType = iota
)

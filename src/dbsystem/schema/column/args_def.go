package column

type Args interface {
	Type() (Type, error)
	Name() (string, error)
	Nullable() (bool, error)
	Length() (uint32, error)
}

func NewSimpleArgs(name string, ctype Type) Args {
	return ArgsBuilder().
		Name(name).
		Type(ctype).
		Build()
}

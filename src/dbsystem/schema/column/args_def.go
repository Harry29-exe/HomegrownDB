package column

type Args interface {
	Type() (Type, error)
	Name() (string, error)
	Nullable() (bool, error)
	Length() (uint32, error)
}

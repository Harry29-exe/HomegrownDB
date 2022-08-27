package parser

type Tree interface {
	Type() RootType
	Root() any
}

type BasicTree struct {
	rootType RootType
	root     any
}

func (b BasicTree) Type() RootType {
	return b.rootType
}

func (b BasicTree) Root() any {
	return b.root
}

type RootType = uint16

const (
	SELECT RootType = iota
	INSERT
	UPDATE
	DELETE
)

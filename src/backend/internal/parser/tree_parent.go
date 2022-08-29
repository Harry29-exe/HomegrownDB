package parser

type Tree struct {
	RootType RootType
	Root     any
}

type RootType = uint16

const (
	Select RootType = iota
	Insert
	Update
	Delete
)

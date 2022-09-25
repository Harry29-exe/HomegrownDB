package analyser

type Tree struct {
	RootType rootType
	Root     any
}

type rootType = uint8

const (
	RootTypeSelect rootType = iota
	RootTypeInsert
)

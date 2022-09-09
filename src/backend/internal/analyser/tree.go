package analyser

type Tree struct {
	RootType rootType
	Ctx      *Ctx
	Root     any
}

type rootType = uint8

const (
	RootTypeSelect rootType = iota
	RootTypeInsert
)

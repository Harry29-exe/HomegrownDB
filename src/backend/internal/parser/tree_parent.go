package parser

type Tree struct {
	RootType RootType
	Root     any
}

type RootType = uint16

const (
	SELECT RootType = iota
	INSERT
	UPDATE
	DELETE
)

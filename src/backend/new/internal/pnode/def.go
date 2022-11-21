package node

type PNode struct {
	Tag        Tag
	StartToken uint
	EndToken   uint
}

type Tag = uint16

const (
	TypeExpr Tag = iota
)

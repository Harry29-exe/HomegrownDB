package node

type Node struct {
	Tag
}

type Tag = uint16

const (
	TypeExpr Tag = iota
)

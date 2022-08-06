package plan

type Node interface {
	Type() nodeType
	Children() []Node
}

type nodeType = uint16

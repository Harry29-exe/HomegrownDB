package plan

type Node interface {
	Type() NodeType
	Children() []Node
}

type NodeType = uint16

const (
	SelectFieldsNode NodeType = iota
	SeqScanNode
)

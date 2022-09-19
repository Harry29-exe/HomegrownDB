package pnode

type TokenIndex = uint32

func NewNode(start, end TokenIndex) Node {
	return Node{
		NodeStartTokenIndex: start,
		NodeEndTokenIndex:   end,
	}
}

type Node struct {
	NodeStartTokenIndex TokenIndex
	NodeEndTokenIndex   TokenIndex
}

func (n *Node) SetTokenIndexes(start, end TokenIndex) {
	n.NodeStartTokenIndex = start
	n.NodeEndTokenIndex = end
}

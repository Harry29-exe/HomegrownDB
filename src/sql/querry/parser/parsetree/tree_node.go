package parsetree

import (
	"errors"
	"strconv"
)

type Node interface {
	Children() []Node
	AddChild(node Node) error
	Value() any
	Type() NodeType
}

type NodeType = uint16

const (
	Select NodeType = iota // Select node has nil value, can be created with NewSelectNode
	From
	Where
	Table
	Fields
	Field
)

type basicNode struct {
	acceptingNodeTypes AcceptedNodes
	children           []Node
	value              any
	nodeType           NodeType
}

type AcceptedNodes = []NodeType

func (n *basicNode) Children() []Node {
	return n.children
}

func (n *basicNode) AddChild(node Node) error {
	accepted := n.canAcceptType(node.Type())
	if !accepted {
		return errors.New("nodes with type: " + strconv.Itoa(int(n.nodeType)) + " does accept nodes of type: " + strconv.Itoa(int(node.Type())))
	}

	n.children = append(n.children, node)
	return nil
}

func (n *basicNode) Value() any {
	return n.value
}

func (n *basicNode) Type() NodeType {
	return n.nodeType
}

func (n *basicNode) canAcceptType(nodeType NodeType) bool {
	for _, node := range n.acceptingNodeTypes {
		if node == nodeType {
			return true
		}
	}
	return false
}

type leafNode struct {
	value    any
	nodeType NodeType
}

func (l *leafNode) Children() []Node {
	return nil
}

func (l *leafNode) AddChild(node Node) error {
	return errors.New("leaf node does not supports adding children")
}

func (l *leafNode) Value() any {
	return l.value
}

func (l *leafNode) Type() NodeType {
	return l.nodeType
}

package node

type Node interface {
	// Tag indicates node type
	Tag() Tag
	// DEqual debug equal, used in tests for fast assertions
	DEqual() bool
}

type node struct {
	tag Tag
}

func (n *node) Tag() Tag {
	return n.tag
}

type Tag = uint16

const (
	TagRTE Tag = iota
	TagRteRef
	TagFrom
	TagAlias

	// TagExprStart expressions nodes start
	TagExprStart
	TagVar
)

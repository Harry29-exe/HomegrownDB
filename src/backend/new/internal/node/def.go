package node

import "reflect"

type Node interface {
	// Tag indicates node type
	Tag() Tag
	// DEqual debug equal, used in tests for fast assertions
	DEqual(node Node) bool
}

type node struct {
	tag Tag
}

func (n *node) Tag() Tag {
	return n.tag
}

type Tag = uint16

const (
	TagQuery Tag = iota
	TagRTE
	TagRteRef
	TagFrom
	TagAlias

	// TagExpr expressions nodes start
	TagExpr
	TagTargetEntry
	TagVar
)

// -------------------------
//      UtilsFunction
// -------------------------

func nodeEqual(v1, v2 Node) (result bool, conclusive bool) {
	if nodesEqNil(v1, v2) {
		return true, true
	} else if !basicNodeEqual(v1, v2) {
		return false, true
	}
	return true, false
}

func nodesEqNil(v1, v2 Node) bool {
	return isNil(v1) && isNil(v2)
}

func basicNodeEqual(v1, v2 Node) bool {
	if isNil(v1) || isNil(v2) {
		return false
	}

	return v1.Tag() == v2.Tag()
}

func isNil(node Node) bool {
	return node == nil || reflect.ValueOf(node).IsNil()
}

func cmpNodeArray[T Node](nodes1, nodes2 []T) bool {
	if len(nodes1) != len(nodes2) {
		return false
	}
	for i := 0; i < len(nodes1); i++ {
		if !nodes1[i].DEqual(nodes2[i]) {
			return false
		}
	}
	return true
}

func cmpNodeArray2D[T Node](nodes1, nodes2 [][]T) bool {
	if len(nodes1) != len(nodes2) {
		return false
	}

	for i := 0; i < len(nodes1); i++ {
		if len(nodes1[i]) != len(nodes2[i]) {
			return false
		}
		for j := 0; j < len(nodes1[i]); j++ {
			if !nodes1[i][j].DEqual(nodes2[i][j]) {
				return false
			}
		}
	}
	return true
}

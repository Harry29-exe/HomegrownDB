package pnode

import "reflect"

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
		if !nodes1[i].Equal(nodes2[i]) {
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
			if !nodes1[i][j].Equal(nodes2[i][j]) {
				return false
			}
		}
	}
	return true
}

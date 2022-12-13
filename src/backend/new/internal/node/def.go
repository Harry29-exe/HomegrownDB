package node

import (
	"fmt"
	"reflect"
	"strings"
)

type Node interface {
	// Tag indicates test type
	Tag() Tag
	dEqual(node Node) bool //todo remove this from node (it exist on DNode)
	DPrint(nesting int) string
}

type node struct {
	tag Tag
}

func (n *node) DEqual(node Node) bool {
	if n == nil && (node == nil) {
		return true
	}
	return false
}

func (n *node) Tag() Tag {
	return n.tag
}

func (n *node) dFormat(innerStr string, nesting int) string {
	b := strings.Builder{}
	for i := 0; i < nesting; i++ {
		b.WriteString("\t")
	}
	nestingStr := b.String()

	return strings.ReplaceAll(n.dTag(nesting)+innerStr, "\n", "\n"+nestingStr)
}

func (n *node) dTag(nesting int) string {
	return fmt.Sprintf("@%s", n.tag.ToString())
}

type Tag uint16

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

	// Plan nodes
	TagPlan
	TagPlanedStmt
	TagModifyTable
	// Scan nodes
	TagScan
	TagSeqScan
)

func (t Tag) ToString() string {
	return [...]string{
		"TagQuery",
		"TagRTE",
		"TagRteRef",
		"TagFrom",
		"TagAlias",
		"TagExpr",
		"TagTargetEntry",
		"TagVar",
		"TagPlan",
		"TagPlanedStmt",
		"TagModifyTable",
		"TagScan",
		"TagSeqScan",
	}[t]
}

// -------------------------
//      UtilsFunction
// -------------------------

func dPrintArr[T Node](nesting int, nodes []T) string {
	builder := strings.Builder{}
	builder.WriteString("[\n")
	for _, dNode := range nodes {
		nodeStr := dNode.DPrint(nesting) + ",\n"
		builder.WriteString(nodeStr)
	}
	builder.WriteString("]\n")
	return builder.String()
}

func DEqual(v1, v2 Node) bool {
	if nodesEqNil(v1, v2) {
		return true
	} else if !basicNodeEqual(v1, v2) {
		return false
	}
	return v1.dEqual(v2)
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
		if !nodes1[i].dEqual(nodes2[i]) {
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
			if !nodes1[i][j].dEqual(nodes2[i][j]) {
				return false
			}
		}
	}
	return true
}

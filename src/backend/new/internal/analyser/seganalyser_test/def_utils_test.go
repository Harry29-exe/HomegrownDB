package seganalyser_test

import (
	"HomegrownDB/backend/new/internal/node"
	"runtime/debug"
	"testing"
)

var Assert = assertNs{}

// assertNs assert namespace
type assertNs struct{}

func (a assertNs) Node(expected node.Node, actual node.Node, t *testing.T) {
	equal := expected.DEqual(actual)
	if equal {
		return
	}

	t.Errorf("expected: \n%+v\nactual: \n%+v\n", expected, actual)
	debug.PrintStack()
}

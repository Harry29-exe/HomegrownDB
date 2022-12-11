package seganalyser_test

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
	"runtime/debug"
	"strings"
	"testing"
)

// -------------------------
//      Assert
// -------------------------

var Assert = assertNs{}

// assertNs assert namespace
type assertNs struct{}

func (a assertNs) Node(expected node.Node, actual node.Node, t *testing.T) {
	equal := node.DEqual(expected, actual)
	if equal {
		return
	}

	t.Errorf("expected: \n%s\nactual: \n%s\n",
		sprintNode(expected),
		sprintNode(actual),
	)
	debug.PrintStack()
}

func sprintNode(node node.Node) string {
	return formatNode(fmt.Sprintf("%#v", node))
}

func formatNode(err string) string {
	err1 := strings.ReplaceAll(err, ", ", ",\n")

	b := strings.Builder{}
	nesting := 0
	for _, c := range err1 {
		switch c {
		case '{':
			nesting++
			b.WriteRune(c)
		case '}':
			nesting--
			b.WriteRune(c)
		case '\n':
			b.WriteRune(c)
			b.WriteString(nestingStr(nesting))
		default:
			b.WriteRune(c)
		}
	}
	return b.String()
}

func nestingStr(nesting int) string {
	return strings.Repeat("\t", nesting)
}

// -------------------------
//      Utils
// -------------------------

func StoreWithUsersTable(t *testing.T) (store table.Store, users table.Definition) {
	users = tt_user.Def(t)
	store = table.NewEmptyTableStore()
	err := store.AddTable(users)
	assert.ErrIsNil(err, t)

	return store, users
}

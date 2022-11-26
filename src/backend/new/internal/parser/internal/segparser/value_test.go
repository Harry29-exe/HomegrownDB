package segparser_test

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/segparser"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestValue_Parse_ColRef1(t *testing.T) {
	query := "u.username"
	expected := pnode.NewColumnRef("username", "u")

	src := internal.NewTokenSource(query)
	v := validator.NewValidator(src)

	node, err := segparser.Value.Parse(src, v)
	assert.IsNil(err, t)
	assert.True(expected.Equal(node), t)
}

package parse_test

import (
	"HomegrownDB/backend/internal/parser/parse"
	"HomegrownDB/backend/internal/parser/tokenizer"
	"HomegrownDB/backend/internal/parser/validator"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

func TestValue_Parse_ColRef1(t *testing.T) {
	query := "u.username"
	expected := pnode.NewColumnRef("username", "u")

	src := tokenizer.NewTokenSource(query)
	v := validator.NewValidator(src)

	node, err := parse.Value.Parse(src, v)
	assert.IsNil(err, t)
	assert.True(expected.Equal(node), t)
}

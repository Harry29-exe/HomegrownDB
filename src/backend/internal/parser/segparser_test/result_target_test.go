package parser

import (
	"HomegrownDB/backend/internal/parser/segparser"
	"HomegrownDB/backend/internal/parser/tokenizer"
	"HomegrownDB/backend/internal/parser/validator"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

// --------------------------------------------------
//      TestResultTarget_Select_Positive
// --------------------------------------------------

func testResultTargetSelect(query string, expectedTree pnode2.Node, t *testing.T) {
	src := tokenizer.NewTokenSource(query)
	v := validator.NewValidator(src)

	resTarget, err := segparser.ResultTarget.Parse(src, v, segparser.ResultTargetSelect)

	assert.ErrIsNil(err, t)
	assert.True(resTarget.Equal(expectedTree), t)
}

func TestResultTarget_Select_Column_Positive1(t *testing.T) {
	query := "u.name"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("name", "u"))
	testResultTargetSelect(query, expectedTree, t)
}

func TestResultTarget_Select_Column_Positive2(t *testing.T) {
	query := "u.name as uname"
	expectedTree := pnode2.NewResultTarget("uname", pnode2.NewColumnRef("name", "u"))
	testResultTargetSelect(query, expectedTree, t)
}
func TestResultTarget_Select_Column_Positive3(t *testing.T) {
	query := "name"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("name", ""))
	testResultTargetSelect(query, expectedTree, t)
}
func TestResultTarget_Select_Column_Positive4(t *testing.T) {
	query := "name as alias_name"
	expectedTree := pnode2.NewResultTarget("alias_name", pnode2.NewColumnRef("name", ""))
	testResultTargetSelect(query, expectedTree, t)
}
func TestResultTarget_Select_Column_Positive5(t *testing.T) {
	query := "u.u"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("u", "u"))
	testResultTargetSelect(query, expectedTree, t)
}
func TestResultTarget_Select_Column_Positive6(t *testing.T) {
	query := "name.u"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("u", "name"))
	testResultTargetSelect(query, expectedTree, t)
}

func TestResultTarget_Select_AConst_Positive1(t *testing.T) {
	query := "'name'"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewAConstStr("name"))

	testResultTargetSelect(query, expectedTree, t)
}

// --------------------------------------------------
//      TestResultTarget_Select_Negative
// --------------------------------------------------

func TestResultTarget_Select_Column_Negative(t *testing.T) {
	queries := []string{
		" u.name",
		" name",
		", name",
		"(ww",
	}

	for _, query := range queries {
		src := tokenizer.NewTokenSource(query)
		v := validator.NewValidator(src)

		res, err := segparser.ResultTarget.Parse(src, v, segparser.ResultTargetSelect)
		_ = res

		assert.ErrNotNil(err, t)
	}
}

// --------------------------------------------------
//      TestResultTarget_Insert_Positive
// --------------------------------------------------

func testResultTargetInsert(query string, expectedTree pnode2.Node, t *testing.T) {
	src := tokenizer.NewTokenSource(query)
	v := validator.NewValidator(src)

	resTarget, err := segparser.ResultTarget.Parse(src, v, segparser.ResultTargetInsert)

	assert.ErrIsNil(err, t)
	assert.True(resTarget.Equal(expectedTree), t)
}

func TestResultTarget_Insert_Column_Positive1(t *testing.T) {
	query := "name"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("name", ""))
	testResultTargetInsert(query, expectedTree, t)
}

func TestResultTarget_Insert_Column_Positive2(t *testing.T) {
	query := "name as uname" // expected to parse name but not 'as uname'
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("name", ""))
	testResultTargetInsert(query, expectedTree, t)
}

func TestResultTarget_Insert_Column_Positive3(t *testing.T) {
	query := "u.name"
	expectedTree := pnode2.NewResultTarget("", pnode2.NewColumnRef("u", ""))
	testResultTargetInsert(query, expectedTree, t)
}

// --------------------------------------------------
//      TestResultTarget_Insert_Negative
// --------------------------------------------------

func TestResultTarget_Insert_Column_Negative(t *testing.T) {
	queries := []string{
		" u.name",
		" name",
		", name",
		"(ww",
		"'name'",
	}

	for _, query := range queries {
		src := tokenizer.NewTokenSource(query)
		v := validator.NewValidator(src)

		res, err := segparser.ResultTarget.Parse(src, v, segparser.ResultTargetInsert)
		_ = res

		assert.ErrNotNil(err, t)
	}
}

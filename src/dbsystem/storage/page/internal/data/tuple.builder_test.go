package data_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	. "HomegrownDB/dbsystem/storage/page/internal/data"
	"HomegrownDB/dbsystem/tx"
	"strings"
	"testing"
)

func TestBuildSimplePage(t *testing.T) {
	// given
	builder, pattern := createBuilderAndPattern()

	// when
	col0 := intype.ConvInt8Value(8)
	err := builder.WriteValue(col0)
	assert.ErrIsNil(err, t)
	col1, err := intype.ConvStrValue("Bob")
	assert.ErrIsNil(err, t)
	err = builder.WriteValue(col1)
	assert.ErrIsNil(err, t)

	tuple := builder.VolatileTuple(tx.StdTx{}, 0)

	// then
	assert.True(pattern.Columns[0].Type.Equal(col0.NormValue, tuple.ColValue(0)), t)
	assert.True(pattern.Columns[1].Type.Equal(col1.NormValue, tuple.ColValue(1)), t)
}

func TestBuildMultiplePages(t *testing.T) {
	// given
	builder, pattern := createBuilderAndPattern()

	// when
	// creating first tuple
	col0v1 := intype.ConvInt8Value(8)
	err := builder.WriteValue(col0v1)
	assert.ErrIsNil(err, t)
	col1v1, err := intype.ConvStrValue("Bob")
	assert.ErrIsNil(err, t)
	err = builder.WriteValue(col1v1)
	assert.ErrIsNil(err, t)
	tupleV1 := builder.VolatileTuple(tx.StdTx{}, 0)

	// starting creating second
	builder.Reset()

	col0v2 := intype.ConvInt8Value(14)
	err = builder.WriteValue(col0v2)
	assert.ErrIsNil(err, t)
	// testing whether value has been overridden (normally we should not use volatileTuple after builder.Reset())
	assert.True(pattern.Columns[0].Type.Equal(col0v2.NormValue, tupleV1.ColValue(0)), t)
	col1v2, err := intype.ConvStrValue(strings.Repeat("Alice", 200))
	assert.ErrIsNil(err, t)
	err = builder.WriteValue(col1v2)
	assert.ErrIsNil(err, t)
	tupleV2 := builder.VolatileTuple(tx.StdTx{}, 0)

	// then
	assert.True(pattern.Columns[0].Type.Equal(col0v2.NormValue, tupleV2.ColValue(0)), t)
	assert.True(pattern.Columns[1].Type.Equal(col1v2.NormValue, tupleV2.ColValue(1)), t)
}

func createBuilderAndPattern() (TupleBuilder, TuplePattern) {
	builder := NewTupleBuilder()
	pattern := NewPattern([]PatternCol{
		{
			Type: hgtype.NewInt8(rawtype.Args{}),
			Name: "col0",
		},
		{
			Type: hgtype.NewStr(rawtype.Args{Length: 10_000, VarLen: true}),
			Name: "col1",
		},
	})
	builder.Init(pattern)
	return builder, pattern
}

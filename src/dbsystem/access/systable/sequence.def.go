package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
)

var sequenceDef = createHGSequences()

func SequencesTableDef() reldef.TableRDefinition {
	return sequenceDef
}

func createHGSequences() reldef.TableRDefinition {
	columns := []reldef.ColumnDefinition{
		sequencesCols.oid(),
		sequencesCols.typeTag(),
		sequencesCols.seqStart(),
		sequencesCols.seqIncrement(),
		sequencesCols.seqMax(),
		sequencesCols.seqMin(),
		sequencesCols.seqCacheSize(),
		sequencesCols.seqCycle(),
	}

	return internal.NewTableDef(
		SequencesName,
		HGSequencesOID,
		HGSequencesFsmOID,
		HGSequencesVmOID,
		columns,
	)
}

const (
	SequencesOrderOID reldef.Order = iota
	SequencesOrderTypeTag
	SequencesOrderSeqStart
	SequencesOrderSeqIncrement
	SequencesOrderSeqMax
	SequencesOrderSeqMin
	SequencesOrderSeqCacheSize
	SequencesOrderSeqCycle
)

const (
	SequencesColNameOID          = "oid"
	SequencesColNameTypeTag      = "type_tag"
	SequencesColNameSeqStart     = "seq_start"
	SequencesColNameSeqIncrement = "seq_increment"
	SequencesColNameSeqMax       = "seq_max"
	SequencesColNameSeqMin       = "seq_min"
	SequencesColNameSeqCacheSize = "seq_cache_size"
	SequencesColNameSeqCycle     = "seq_cycle"
)

var sequencesCols = sequencesColBuilder{}

type sequencesColBuilder struct{}

func (s sequencesColBuilder) oid() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameOID,
		HGSequencesColOID,
		SequencesOrderOID,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) typeTag() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameTypeTag,
		HGSequencesColTypeTag,
		SequencesOrderTypeTag,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqStart() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqStart,
		HGSequencesColSeqStart,
		SequencesOrderSeqStart,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqIncrement() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqIncrement,
		HGSequencesColSeqIncrement,
		SequencesOrderSeqIncrement,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMax() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqMax,
		HGSequencesColSeqMax,
		SequencesOrderSeqMax,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMin() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqMin,
		HGSequencesColSeqMin,
		SequencesOrderSeqMin,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

// seqCacheSize is not used at this moment, it was created as
// postgresql has this but for simplicity is not used for now
func (sequencesColBuilder) seqCacheSize() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqCacheSize,
		HGSequencesColSeqCacheSize,
		SequencesOrderSeqCacheSize,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqCycle() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		SequencesColNameSeqCycle,
		HGSequencesColSeqCycle,
		SequencesOrderSeqCycle,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

var sequenceDef = createHGSequences()

func SequencesTableDef() tabdef.TableRDefinition {
	return sequenceDef
}

func createHGSequences() tabdef.TableRDefinition {
	columns := []tabdef.ColumnDefinition{
		sequencesCols.oid(),
		sequencesCols.typeTag(),
		sequencesCols.seqStart(),
		sequencesCols.seqIncrement(),
		sequencesCols.seqMax(),
		sequencesCols.seqMin(),
		sequencesCols.seqCache(),
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
	SequencesOrderOID tabdef.Order = iota
	SequencesOrderTypeTag
	SequencesOrderSeqStart
	SequencesOrderSeqIncrement
	SequencesOrderSeqMax
	SequencesOrderSeqMin
	SequencesOrderSeqCache
	SequencesOrderSeqCycle
)

const (
	SequencesColNameOID          = "oid"
	SequencesColNameTypeTag      = "type_tag"
	SequencesColNameSeqStart     = "seq_start"
	SequencesColNameSeqIncrement = "seq_increment"
	SequencesColNameSeqMax       = "seq_max"
	SequencesColNameSeqMin       = "seq_min"
	SequencesColNameSeqCache     = "seq_cache"
	SequencesColNameSeqCycle     = "seq_cycle"
)

var sequencesCols = sequencesColBuilder{}

type sequencesColBuilder struct{}

func (s sequencesColBuilder) oid() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameOID,
		HGSequencesColOID,
		SequencesOrderOID,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) typeTag() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameTypeTag,
		HGSequencesColTypeTag,
		SequencesOrderTypeTag,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqStart() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqStart,
		HGSequencesColSeqStart,
		SequencesOrderSeqStart,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqIncrement() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqIncrement,
		HGSequencesColSeqIncrement,
		SequencesOrderSeqIncrement,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMax() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqMax,
		HGSequencesColSeqMax,
		SequencesOrderSeqMax,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMin() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqMin,
		HGSequencesColSeqMin,
		SequencesOrderSeqMin,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqCache() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqCache,
		HGSequencesColSeqCache,
		SequencesOrderSeqCache,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqCycle() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		SequencesColNameSeqCycle,
		HGSequencesColSeqCycle,
		SequencesOrderSeqCycle,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

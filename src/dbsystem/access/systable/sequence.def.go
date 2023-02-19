package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

var sequenceDef = createHGSequences()

func SequencesTableDef() tabdef.RDefinition {
	return sequenceDef
}

func createHGSequences() tabdef.RDefinition {
	columns := []column.WDef{
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
	SequencesOrderOID column.Order = iota
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

func (s sequencesColBuilder) oid() column.WDef {
	return column.NewDefinition(
		SequencesColNameOID,
		HGSequencesColOID,
		SequencesOrderOID,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) typeTag() column.WDef {
	return column.NewDefinition(
		SequencesColNameTypeTag,
		HGSequencesColTypeTag,
		SequencesOrderTypeTag,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqStart() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqStart,
		HGSequencesColSeqStart,
		SequencesOrderSeqStart,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqIncrement() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqIncrement,
		HGSequencesColSeqIncrement,
		SequencesOrderSeqIncrement,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMax() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqMax,
		HGSequencesColSeqMax,
		SequencesOrderSeqMax,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqMin() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqMin,
		HGSequencesColSeqMin,
		SequencesOrderSeqMin,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqCache() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqCache,
		HGSequencesColSeqCache,
		SequencesOrderSeqCache,
		hgtype.NewInt8(hgtype.Args{Nullable: false}),
	)
}

func (sequencesColBuilder) seqCycle() column.WDef {
	return column.NewDefinition(
		SequencesColNameSeqCycle,
		HGSequencesColSeqCycle,
		SequencesOrderSeqCycle,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

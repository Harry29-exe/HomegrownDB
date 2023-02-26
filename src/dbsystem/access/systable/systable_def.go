package systable

import (
	"HomegrownDB/dbsystem/hglib"
)

type OID = hglib.OID

const (
	RelationsName = "sys_relations"
	ColumnsName   = "sys_columns"
	SequencesName = "sys_sequences"
)

const (
	_ OID = iota
	HGRelationsOID
	HGRelationsFsmOID
	HGRelationsVmOID

	HGRelationsColOID
	HGRelationsColRelKind
	HGRelationsColRelName
	HGRelationsColFsmOID
	HGRelationsColVmOID

	// HGColumnsOID start of sys_columns

	HGColumnsOID
	HGColumnsFsmOID
	HGColumnsVmOID

	HGColumnsColOID
	HGColumnsColRelationOID
	HGColumnsColColName
	HGColumnsColColOrder
	HGColumnsColTypeTag
	HGColumnsColArgsLength
	HGColumnsColArgsNullable
	HGColumnsColArgsVarLen
	HGColumnsColArgsUTF8

	//HGSequencesOID start of sys_sequences

	HGSequencesOID
	HGSequencesFsmOID
	HGSequencesVmOID

	HGSequencesColOID
	HGSequencesColTypeTag
	HGSequencesColSeqStart
	HGSequencesColSeqIncrement
	HGSequencesColSeqMax
	HGSequencesColSeqMin
	HGSequencesColSeqCacheSize
	HGSequencesColSeqCycle

	HGObjectsOIDsStart
)

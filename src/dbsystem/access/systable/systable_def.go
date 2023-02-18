package systable

import (
	"HomegrownDB/dbsystem/hglib"
)

type OID = hglib.OID

const (
	RelationsName = "sys_relations"
	ColumnsName   = "sys_columns"
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
)

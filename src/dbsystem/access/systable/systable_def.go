package systable

import "HomegrownDB/dbsystem/relation/dbobj"

type OID = dbobj.OID

const (
	RelationsName = "sys_relations"
	ColumnsName   = "sys_columns"
)

const (
	HGRelationsOID OID = iota
	HGRelationsFsmOID
	HGRelationsVmOID

	HGRelationsColOID
	HGRelationsColRelKind
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

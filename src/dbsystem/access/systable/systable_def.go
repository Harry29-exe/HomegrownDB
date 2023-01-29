package systable

import "HomegrownDB/dbsystem/relation/dbobj"

type OID = dbobj.OID

const (
	RelationsName = "sys_relations"
	ColumnsName   = "sys_columns"
)

const (
	RelationsOID OID = iota
	RelationsFsmOID
	RelationsVmOID
	ColumnsOID
	ColumnsFsmOID
	ColumnsVmOID
)

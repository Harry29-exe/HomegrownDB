package systable

import "HomegrownDB/dbsystem/relation/dbobj"

const (
	RelationsName = "sys_relations"
)

const (
	RelationsOID dbobj.OID = iota
	RelationsFsmOID
	RelationsVmOID
)

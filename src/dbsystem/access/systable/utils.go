package systable

import (
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
	"log"
)

func newTableDef(name string, oid OID, fsmOID OID, vmOID OID, columns []column.WDef) table.RDefinition {
	tableDef := table.NewDefinition(name)
	tableDef.InitRel(oid, fsmOID, vmOID)

	for _, colDef := range columns {
		err := tableDef.AddColumn(colDef)
		if err != nil {
			log.Panicf("unexpected error while creating Relations table: %s", err.Error())
		}
	}

	return tableDef
}

package internal

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
	"log"
)

func NewTableDef(name string, oid reldef.OID, fsmOID reldef.OID, vmOID reldef.OID, columns []reldef.ColumnDefinition) reldef.TableRDefinition {
	tableDef := reldef.NewTableDefinition(name)
	tableDef.InitRel(oid, fsmOID, vmOID)

	for _, colDef := range columns {
		err := tableDef.AddColumn(colDef)
		if err != nil {
			log.Panicf("unexpected error while creating Relations tabdef: %s", err.Error())
		}
	}

	return tableDef
}

func NewColumnDef(name string, oid reldef.OID, order reldef.Order, ctype hgtype.ColType) reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(name, oid, order, ctype)
}

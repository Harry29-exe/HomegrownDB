package sysinit

import (
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/tx"
)

func CreateSysTables(fs dbfs.FS) error {
	creator := sysTablesCreator{FS: fs}
	relationsTable := systable.RelationsTableDef()
	columnsTable := systable.ColumnsTableDef()
	sequencesTable := systable.SequencesTableDef()
	creator.cache = newPageCache(creator.FS, relationsTable, columnsTable)

	creatorTX := tx.StdTx{Id: 0}
	return creator.
		createTable(systable.HGRelationsOID, systable.HGRelationsFsmOID, systable.HGRelationsVmOID).
		createTable(systable.HGColumnsOID, systable.HGColumnsFsmOID, systable.HGColumnsVmOID).
		createTable(systable.HGSequencesOID, systable.HGSequencesFsmOID, systable.HGSequencesVmOID).
		insertTuples(relationsTable.OID(),
			panicOnErr(systable.RelationsOps.TableAsRelationsRow(relationsTable, creatorTX)),
			panicOnErr(systable.RelationsOps.TableAsRelationsRow(columnsTable, creatorTX)),
			panicOnErr(systable.RelationsOps.TableAsRelationsRow(sequencesTable, creatorTX)),
		).
		insertTuplesArr(systable.HGColumnsOID,
			systable.ColumnsOps.DataToRows(systable.HGRelationsOID, relationsTable.Columns(), creatorTX),
			systable.ColumnsOps.DataToRows(systable.HGColumnsOID, columnsTable.Columns(), creatorTX),
			systable.ColumnsOps.DataToRows(systable.HGSequencesOID, sequencesTable.Columns(), creatorTX),
		).
		flushPages().
		getError()
}

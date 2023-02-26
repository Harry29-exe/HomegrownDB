package sysinit

import (
	"HomegrownDB/dbsystem/access/sysinit/internal"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/tx"
)

var Phase1 = phase1{}

func (phase1) Execute(fs dbfs.FS) error {
	return createSysTables(fs)
}

func createSysTables(fs dbfs.FS) error {
	creator := internal.SysTablesCreator{FS: fs}
	relationsTable := systable.RelationsTableDef()
	columnsTable := systable.ColumnsTableDef()
	sequencesTable := systable.SequencesTableDef()
	creator.Cache = internal.NewPageCache(creator.FS, relationsTable, columnsTable)

	creatorTX := tx.StdTx{Id: 0}
	return creator.
		CreateTable(systable.HGRelationsOID, systable.HGRelationsFsmOID, systable.HGRelationsVmOID).
		CreateTable(systable.HGColumnsOID, systable.HGColumnsFsmOID, systable.HGColumnsVmOID).
		CreateTable(systable.HGSequencesOID, systable.HGSequencesFsmOID, systable.HGSequencesVmOID).
		InsertTuples(relationsTable.OID(),
			internal.PanicOnErr(systable.RelationsOps.TableAsRelationsRow(relationsTable, creatorTX)),
			internal.PanicOnErr(systable.RelationsOps.TableAsRelationsRow(columnsTable, creatorTX)),
			internal.PanicOnErr(systable.RelationsOps.TableAsRelationsRow(sequencesTable, creatorTX)),
		).
		InsertTuplesArr(systable.HGColumnsOID,
			systable.ColumnsOps.DataToRows(systable.HGRelationsOID, relationsTable.Columns(), creatorTX),
			systable.ColumnsOps.DataToRows(systable.HGColumnsOID, columnsTable.Columns(), creatorTX),
			systable.ColumnsOps.DataToRows(systable.HGSequencesOID, sequencesTable.Columns(), creatorTX),
		).
		FlushPages().
		GetError()
}

type phase1 struct{}

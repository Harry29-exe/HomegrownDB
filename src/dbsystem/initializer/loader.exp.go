package initializer

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
)

func ExpLoadDB() (dbsystem.DBSystem, error) {
	var err error
	ctx := &LoaderCtx{}

	ctx.thenAll(func() {
		ctx.RootPath, ctx.err = config.ReadRootPath()
	}, func() {
		ctx.FS = &dbfs.StdFS{RootPath: ctx.RootPath}
	}, func() {
		ctx.Config, ctx.err = config.ReadConfig(ctx.FS)
	}, func() {
		ctx.InitProps, ctx.err = readInitProperties(ctx.FS)
	}, func() {
		ctx.PageIOStore = pageio.NewStore(ctx.FS)
	}, func() {
		ctx.SharedBuffer = buffer.NewSharedBuffer(ctx.Config.SharedBufferSize, ctx.PageIOStore)
	}, func() {
		ctx.FsmStore = fsm.NewStore()
	}, func() {
		ctx.TableStore = table.NewEmptyTableStore()
	}).thenValidate().
		after(exprLoadRelations)

	system := &dbsystem.StdSystem{
		PageIO:   ctx.PageIOStore,
		Tables:   ctx.TableStore,
		FSMs:     ctx.FsmStore,
		DBBuffer: ctx.SharedBuffer,
	}

	err = system.SetState(dbsystem.State{
		NextRID: ctx.InitProps.NextRID,
		NextOID: ctx.InitProps.NextOID,
	})
	return system, err
}

func exprLoadRelations(ctx *LoaderCtx) error {
	var err error
	for _, rel := range ctx.InitProps.Relations {
		var fileData []byte
		fileData, err = readRelationFile(rel.RelationID, ctx)
		if err != nil {
			return err
		}

		switch rel.RelKind {
		case relation.TypeTable:
			err = loadTable(fileData, ctx)
		case relation.TypeFsm:
			err = loadFSM(fileData, ctx)
		default:
			//todo implement me
			panic("Not implemented")
		}
	}

	return err
}

func exprLoadTable(serializedData []byte, ctx *LoaderCtx) error {
	tableDef := table.Deserialize(serializedData)
	err := ctx.TableStore.LoadTable(tableDef)
	if err != nil {
		return err
	}
	return openFileAndRegisterRelationIO(tableDef, ctx)
}

func exprLoadFSM(serializedData []byte, ctx *LoaderCtx) error {
	deserializer := bparse.NewDeserializer(serializedData)
	fsmRel := fsm.DeserializeFSM(ctx.SharedBuffer, deserializer)
	ctx.FsmStore.RegisterFSM(fsmRel)
	return openFileAndRegisterRelationIO(fsmRel, ctx)
}

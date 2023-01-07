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

type loadFunc func(loadCtx *LoaderCtx) (err error)

var loadSteps = []loadFunc{
	loadRootPath,
	loadDBFS,
	loadConfig,
	loadProps,
	loadPageIOStore,
	loadBuffer,
	loadFsmStore,
	loadTableStore,
	loadRelations,
}

func LoadDB() (dbsystem.DBSystem, error) {
	var err error
	ctx := &LoaderCtx{}

	for _, step := range loadSteps {
		err = step(ctx)
		if err != nil {
			return nil, err
		}
	}

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

func loadRootPath(ctx *LoaderCtx) error {
	var err error
	ctx.RootPath, err = config.ReadRootPath()
	return err
}

func loadConfig(ctx *LoaderCtx) error {
	var err error
	ctx.Config, err = config.ReadConfig(ctx.FS)
	return err
}

func loadProps(ctx *LoaderCtx) error {
	var err error
	ctx.InitProps, err = readInitProperties(ctx.FS)
	return err
}

func loadDBFS(ctx *LoaderCtx) error {
	ctx.FS = &dbfs.StdFS{RootPath: ctx.Config.DBHomePath}
	return nil
}

func loadPageIOStore(ctx *LoaderCtx) error {
	ctx.PageIOStore = pageio.NewStore(ctx.FS)
	return nil
}

func loadBuffer(ctx *LoaderCtx) error {
	buff := buffer.NewSharedBuffer(ctx.Config.SharedBufferSize, ctx.PageIOStore)
	ctx.SharedBuffer = buff
	return nil
}

func loadFsmStore(initCtx *LoaderCtx) error {
	initCtx.FsmStore = fsm.NewStore()
	return nil
}

func loadTableStore(initCtx *LoaderCtx) error {
	initCtx.TableStore = table.NewEmptyTableStore()
	return nil
}

func loadRelations(ctx *LoaderCtx) error {
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

func loadTable(serializedData []byte, ctx *LoaderCtx) error {
	tableDef := table.Deserialize(serializedData)
	err := ctx.TableStore.LoadTable(tableDef)
	if err != nil {
		return err
	}
	return openFileAndRegisterRelationIO(tableDef, ctx)
}

func loadFSM(serializedData []byte, ctx *LoaderCtx) error {
	deserializer := bparse.NewDeserializer(serializedData)
	fsmRel := fsm.DeserializeFSM(ctx.SharedBuffer, deserializer)
	ctx.FsmStore.RegisterFSM(fsmRel)
	return openFileAndRegisterRelationIO(fsmRel, ctx)
}

// -------------------------
//      helpers
// -------------------------

func openFileAndRegisterRelationIO(relation relation.Relation, ctx *LoaderCtx) error {
	file, err := ctx.FS.OpenRelationDataFile(relation)
	if err != nil {
		return err
	}
	pageIO, err := pageio.NewPageIO(file)
	if err != nil {
		return err
	}
	ctx.PageIOStore.Register(relation.RelationID(), pageIO)
	return nil
}

func readRelationFile(id relation.ID, ctx *LoaderCtx) ([]byte, error) {
	file, err := ctx.FS.OpenRelationDef(id)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	buff := make([]byte, stat.Size())
	_, err = file.Read(buff)
	return buff, err
}

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
	"os"
)

type initFunc func(initCtx *ctx) error

var initSteps = []initFunc{
	initDBFS,
	initPageIOStore,
	initBuffer,
	initFsmStore,
	initTableStore,
	initRelations,
}

func InitializeDB(props *config.Properties) (dbsystem.DBSystem, error) {
	initFile, err := readInitProperties()
	if err != nil {
		return nil, err
	}
	ctx := createCtx(props, initFile)

	for _, step := range initSteps {
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

func initDBFS(ctx *ctx) error {
	ctx.FS = &dbfs.StdFS{Rootpath: ctx.Props.DBHomePath}
	return nil
}

func initPageIOStore(ctx *ctx) error {
	ctx.PageIOStore = pageio.NewStore(ctx.FS)
	return nil
}

func initBuffer(initCtx *ctx) error {
	buff := buffer.NewSharedBuffer(initCtx.Props.SharedBufferSize, initCtx.PageIOStore)
	initCtx.SharedBuffer = buff
	return nil
}

func initFsmStore(initCtx *ctx) error {
	initCtx.FsmStore = fsm.NewStore()
	return nil
}

func initTableStore(initCtx *ctx) error {
	initCtx.TableStore = table.NewEmptyTableStore()
	return nil
}

func initRelations(ctx *ctx) error {
	var err error
	for _, rel := range ctx.InitProps.Relations {
		var fileData []byte
		fileData, err = readRelationFile(rel.RelationID, ctx)
		if err != nil {
			return err
		}

		switch rel.RelKind {
		case relation.TypeTable:
			err = initTable(fileData, ctx)
		case relation.TypeFsm:
			err = initFSM(fileData, ctx)
		default:
			//todo implement me
			panic("Not implemented")
		}
	}

	return err
}

func initTable(serializedData []byte, ctx *ctx) error {
	tableDef := table.Deserialize(serializedData)
	err := ctx.TableStore.LoadTable(tableDef)
	if err != nil {
		return err
	}
	return openFileAndRegisterRelationIO(tableDef, ctx)
}

func initFSM(serializedData []byte, ctx *ctx) error {
	deserializer := bparse.NewDeserializer(serializedData)
	fsmRel := fsm.DeserializeFSM(ctx.SharedBuffer, deserializer)
	ctx.FsmStore.RegisterFSM(fsmRel)
	return openFileAndRegisterRelationIO(fsmRel, ctx)
}

// -------------------------
//      helpers
// -------------------------

func openFileAndRegisterRelationIO(relation relation.Relation, ctx *ctx) error {
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

func readRelationFile(id relation.ID, ctx *ctx) ([]byte, error) {
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

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, s.Size())
	_, err = f.Read(data)
	return data, err
}

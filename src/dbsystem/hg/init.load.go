package hg

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"errors"
)

func LoadDB() (DB, error) {
	var err error
	ctx := &loaderCtx{}

	ctx.thenAll(func() {
		ctx.RootPath, ctx.err = config.ReadRootPathEnv()
	}, func() {
		ctx.FS = &dbfs.StdFS{RootPath: ctx.RootPath}
	}, func() {
		ctx.Config, ctx.err = config.ReadConfig(ctx.FS)
	}, func() {
		ctx.DBProps, ctx.err = config.ReadInitProperties(ctx.FS)
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

	system := &DBSystem{
		PageIO:   ctx.PageIOStore,
		Tables:   ctx.TableStore,
		FSMs:     ctx.FsmStore,
		DBBuffer: ctx.SharedBuffer,
	}

	err = system.SetState(State{
		NextRID: ctx.DBProps.NextRID,
		NextOID: ctx.DBProps.NextOID,
	})
	return system, err
}

func exprLoadRelations(ctx *loaderCtx, db DB) error {
	var err error
	for _, rel := range ctx.DBProps.Relations {
		var fileData []byte
		fileData, err = readRelationFile(rel.RelationID, ctx)
		if err != nil {
			return err
		}

		switch rel.RelKind {
		case relation.TypeTable:
			err = exprLoadTable(fileData, ctx)
		case relation.TypeFsm:
			err = exprLoadFSM(fileData, ctx)
		default:
			//todo implement me
			panic("Not implemented")
		}
	}

	return err
}

func exprLoadTable(serializedData []byte, ctx *loaderCtx) error {
	tableDef := table.Deserialize(serializedData)
	err := ctx.TableStore.LoadTable(tableDef)
	if err != nil {
		return err
	}
	return openFileAndRegisterRelationIO(tableDef, ctx)
}

func exprLoadFSM(serializedData []byte, ctx *loaderCtx) error {
	deserializer := bparse.NewDeserializer(serializedData)
	fsmRel := fsm.DeserializeFSM(ctx.SharedBuffer, deserializer)
	ctx.FsmStore.RegisterFSM(fsmRel)
	return openFileAndRegisterRelationIO(fsmRel, ctx)
}

// -------------------------
//      loaderCtx
// -------------------------

type ConfiFun = func(ctx *loaderCtx, db DB) error

type loaderCtx struct {
	DB DB

	RootPath     string
	FS           dbfs.FS
	Config       *config.Configuration
	DBProps      config.DBProperties
	PageIOStore  *pageio.StdStore
	SharedBuffer buffer.SharedBuffer
	TableStore   table.Store
	FsmStore     fsm.Store

	err error
}

func (c *loaderCtx) then(expr func()) *loaderCtx {
	if c.err != nil {
		return c
	}
	expr()
	return c
}

func (c *loaderCtx) thenAll(expressions ...func()) *loaderCtx {
	if c.err != nil {
		return c
	}
	for _, expr := range expressions {
		expr()
		if c.err != nil {
			return c
		}
	}
	return c
}

func (c *loaderCtx) thenValidate() *loaderCtx {
	loader := loaderCtx{}
	if c.RootPath == loader.RootPath ||
		c.FS == nil ||
		c.Config == nil ||
		c.PageIOStore == nil ||
		c.SharedBuffer == nil ||
		c.TableStore == nil ||
		c.FsmStore == nil {
		c.err = errors.New("ctx is invalid")
	}
	return c
}

func (c *loaderCtx) thenCreate() *loaderCtx {
	if c.err != nil {
		return c
	}
	c.DB = &DBSystem{
		Tables:   c.TableStore,
		FSMs:     c.FsmStore,
		PageIO:   c.PageIOStore,
		DBBuffer: c.SharedBuffer,
		FS:       c.FS,
	}
	return c
}

func (c *loaderCtx) after(fn ConfiFun) *loaderCtx {
	if c.err != nil {
		return c
	} else if c.DB == nil {
		c.err = errors.New("during configuration phase loaderCtx.DB can not be null")
		return c
	}
	c.err = fn(c, c.DB)
	return c
}

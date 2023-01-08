package initializer

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/hg"
	table2 "HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"errors"
)

func createCtx(props *config.Configuration, initProps DBProperties) *LoaderCtx {
	return &LoaderCtx{
		Config:    props,
		InitProps: initProps,
	}
}

type LoaderCtx struct {
	RootPath string
	FS       dbfs.FS

	Config    *config.Configuration
	InitProps DBProperties

	PageIOStore  *pageio.StdStore
	SharedBuffer buffer.SharedBuffer
	TableStore   table2.Store
	FsmStore     fsm.Store

	err error
}

func (c *LoaderCtx) BuildDB() (hg.DBStore, error) {
	system := &hg.DBSystem{
		Tables:   c.TableStore,
		FSMs:     c.FsmStore,
		PageIO:   c.PageIOStore,
		DBBuffer: c.SharedBuffer,
	}
	err := system.SetState(hg.State{
		NextRID: c.InitProps.NextRID,
		NextOID: c.InitProps.NextOID,
	})
	if err != nil {
		return nil, err
	}

	return system, nil
}

func (c *LoaderCtx) then(expr func()) *LoaderCtx {
	if c.err != nil {
		return c
	}
	expr()
	return c
}

func (c *LoaderCtx) thenAll(expressions ...func()) *LoaderCtx {
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

func (c *LoaderCtx) thenValidate() *LoaderCtx {
	loader := LoaderCtx{}
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

func (c *LoaderCtx) after(expr func(ctx *LoaderCtx) error) *LoaderCtx {
	if c.err != nil {
		return c
	}
	c.err = expr(c)
	return c
}

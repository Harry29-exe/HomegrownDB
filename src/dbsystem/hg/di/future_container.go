package di

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"fmt"
)

type SimpleArgs struct {
	C  *config.Configuration
	P  config.DBProperties
	FS dbfs.FS
}

type (
	RootProvider         = func() (string, error)
	FsProvider           = func(rootPath string) (dbfs.FS, error)
	ConfigProvider       = func(fs dbfs.FS) (*config.Configuration, error)
	PropertiesProvider   = func(fs dbfs.FS) (config.DBProperties, error)
	PageIOStoreProvider  = func(args SimpleArgs) (pageio.Store, error)
	TableStoreProvider   = func(args SimpleArgs) (table.Store, error)
	FsmStoreProvider     = func(args SimpleArgs) (fsm.Store, error)
	SharedBufferProvider = func(args SimpleArgs, store pageio.Store) (buffer.SharedBuffer, error)
)

type FutureContainer struct {
	RootProvider         RootProvider
	FsProvider           FsProvider
	ConfigProvider       ConfigProvider
	PropertiesProvider   PropertiesProvider
	PageIOStoreProvider  PageIOStoreProvider
	TableStoreProvider   TableStoreProvider
	FsmStoreProvider     FsmStoreProvider
	SharedBufferProvider SharedBufferProvider

	container *Container
	err       error
}

func (c *FutureContainer) Build() (*Container, error) {
	container := &Container{}
	c.container = container
	c.validate().thenAll(func() {
		container.RootPath, c.err = c.RootProvider()
	}, func() {
		container.FS, c.err = c.FsProvider(container.RootPath)
	}, func() {
		container.Config, c.err = c.ConfigProvider(container.FS)
	}, func() {
		container.DBProps, c.err = c.PropertiesProvider(container.FS)
	}, func() {
		container.PageIOStore, c.err = c.PageIOStoreProvider(c.args())
	}, func() {
		container.TableStore, c.err = c.TableStoreProvider(c.args())
	}, func() {
		container.FsmStore, c.err = c.FsmStoreProvider(c.args())
	}, func() {
		container.SharedBuffer, c.err = c.SharedBufferProvider(c.args(), container.PageIOStore)
	})

	return container, c.err
}

func (c *FutureContainer) then(expr func()) *FutureContainer {
	if c.err != nil {
		return c
	}
	expr()
	return c
}

func (c *FutureContainer) thenAll(expressions ...func()) *FutureContainer {
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

func (c *FutureContainer) args() SimpleArgs {
	return SimpleArgs{C: c.container.Config, P: c.container.DBProps, FS: c.container.FS}
}

func (c *FutureContainer) validate() *FutureContainer {
	if c.RootProvider == nil ||
		c.FsProvider == nil ||
		c.ConfigProvider == nil ||
		c.PropertiesProvider == nil ||
		c.PageIOStoreProvider == nil ||
		c.TableStoreProvider == nil ||
		c.FsmStoreProvider == nil ||
		c.SharedBufferProvider == nil {
		c.err = fmt.Errorf("future container has missing dependency.\n%+v", c)
	}
	return nil
}

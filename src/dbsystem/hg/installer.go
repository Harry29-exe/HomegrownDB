package hg

import (
	"HomegrownDB/dbsystem/access/systable/sysinit"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/storage"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func CreateDB(args CreateArgs) error {
	ctx := creatorCtx{Props: args}
	ctx.initRootPath().
		initConfigurationAndProperties().
		initDBFilesystem()
	if ctx.err != nil {
		return ctx.err
	}

	err := sysinit.CreateSysTables(ctx.FS)
	if err != nil {
		return err
	}
	return nil
}

// -------------------------
//      CreateArgs
// -------------------------

type Mode uint8

const (
	InstallerModeDB Mode = iota
	InstallerModeTest
)

type CreateArgs struct {
	Mode     Mode
	RootPath string // RootPath path where db will be initialized (nullable)
	Config   config.Configuration
	Props    config.DBProperties
}

// -------------------------
//      Creator
// -------------------------

type creatorCtx struct {
	Props CreateArgs

	RootPath string
	Config   config.Configuration
	DBProps  config.DBProperties
	FS       dbfs.FS

	err error
}

// -------------------------
//      Init
// -------------------------

func (c *creatorCtx) initRootPath() *creatorCtx {
	c.RootPath = c.Props.RootPath
	if c.modeEqDBInstaller() {
		err := storage.SetRootPathEnv(c.RootPath)
		if err != nil {
			return c.error(err)
		}
	}
	return c
}

func (c *creatorCtx) initConfigurationAndProperties() *creatorCtx {
	c.Config = config.DefaultConfiguration(c.RootPath)
	c.DBProps = config.DefaultDBProperties()
	return c
}

func (c *creatorCtx) initDBFilesystem() *creatorCtx {
	fs, err := dbfs.CreateFS(c.RootPath)
	if err != nil {
		return c.error(err)
	}
	c.FS = fs
	err = fs.InitDBSystemDirs()
	if err != nil {
		return c.error(err)
	}

	err = fs.InitDBSystemConfigAndProps(
		config.SerializeConfig(c.Config),
		config.SerializeProperties(c.DBProps))
	if err != nil {
		return c.error(err)
	}

	return c
}

func (c *creatorCtx) modeEqTest() bool {
	return c.Props.Mode == InstallerModeTest
}

func (c *creatorCtx) modeEqDBInstaller() bool {
	return c.Props.Mode == InstallerModeDB
}

func (c *creatorCtx) error(err error) *creatorCtx {
	c.err = err
	return c
}

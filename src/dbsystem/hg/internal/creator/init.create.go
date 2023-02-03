package creator

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/storage/dbfs"
)

func CreateDB(props Props) error {
	ctx := creatorCtx{Props: props}
	ctx.initRootPath().
		initConfigurationAndProperties().
		initDBFilesystem()
	if ctx.err != nil {
		return ctx.err
	}

	err := createSysTables(ctx.FS)
	if err != nil {
		return err
	}
	return nil
}

// -------------------------
//      Props
// -------------------------

type Mode uint8

const (
	DBInstaller Mode = iota
	Test
)

type Props struct {
	Mode     Mode
	RootPath string // RootPath path where db will be initialized (nullable)
	Config   config.Configuration
	Props    config.DBProperties
}

// -------------------------
//      Creator
// -------------------------

type creatorCtx struct {
	Props Props

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
		err := config.SetRootPathEnv(c.RootPath)
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
	return c.Props.Mode == Test
}

func (c *creatorCtx) modeEqDBInstaller() bool {
	return c.Props.Mode == DBInstaller
}

func (c *creatorCtx) error(err error) *creatorCtx {
	c.err = err
	return c
}

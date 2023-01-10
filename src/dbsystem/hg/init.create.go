package hg

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
	"os"
)

func CreateDB(props CreatorProps) error {
	ctx := creatorCtx{Props: props}
	ctx.initRootPath().
		initConfigurationAndProperties().
		initDBFilesystem().
		createRelations()
	return ctx.err
}

// -------------------------
//      Props
// -------------------------

type CreatorMode uint8

const (
	DBInstaller CreatorMode = iota
	Test
)

type CreatorProps struct {
	Mode              CreatorMode
	RootPath          string              // RootPath path where db will be initialized (nullable)
	RelationsToCreate []relation.Relation // Relations that DBCreator will create
	Config            config.Configuration
	Props             config.DBProperties
}

func (c *CreatorProps) initEmptyWithDefault() error {
	if c.RootPath == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		c.RootPath = homedir + "/.HomegrownDB"
	}
	if c.RelationsToCreate == nil {

	}

	return nil
}

// -------------------------
//      Creator
// -------------------------

type creatorCtx struct {
	Props CreatorProps

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
	err = fs.InitDBSystem()
	if err != nil {
		return c.error(err)
	}
	c.FS = fs
	return c
}

// -------------------------
//      Create
// -------------------------

func (c *creatorCtx) createRelations() *creatorCtx {
	if c.err != nil {
		return c
	}
	relations := c.Props.RelationsToCreate
	for _, relation := range relations {
		err := c.FS.InitNewRelationDir(relation.RelationID())
		if err != nil {
			return c.error(err)
		}
		defFile, err := c.FS.OpenRelationDef(relation.RelationID())
		if err != nil {
			return c.error(err)
		}
		_, err = defFile.Write(SerializeRel(relation))
		if err != nil {
			return c.error(err)
		}
		err = defFile.Close()
		if err != nil {
			return c.error(err)
		}
	}
	return nil
}

// -------------------------
//      Internal
// -------------------------

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

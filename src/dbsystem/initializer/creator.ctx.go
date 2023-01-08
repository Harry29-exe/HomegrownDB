package initializer

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type CreatorCtx struct {
	Props CreatorProps

	RootPath string
	Config   config.Configuration
	DBProps  DBProperties
	FS       dbfs.FS

	err error
}

// -------------------------
//      Init
// -------------------------

func (c *CreatorCtx) initRootPath() *CreatorCtx {
	c.RootPath = c.Props.RootPath
	if c.modeEqDBInstaller() {
		err := config.SetRootPathEnv(c.RootPath)
		if err != nil {
			return c.error(err)
		}
	}
	return c
}

func (c *CreatorCtx) initConfigurationAndProperties() *CreatorCtx {
	c.Config = config.DefaultConfiguration(c.RootPath)
	c.DBProps = DefaultDBProperties()
	return c
}

func (c *CreatorCtx) initDBFilesystem() *CreatorCtx {
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

func (c *CreatorCtx) createRelations() *CreatorCtx {
	if c.err != nil {
		return c
	}
	relations := c.Props.RelationsToCreate
	for _, relation := range relations {
		err := c.FS.CreateRelationDir(relation)
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

func (c *CreatorCtx) modeEqTest() bool {
	return c.Props.Mode == Test
}

func (c *CreatorCtx) modeEqDBInstaller() bool {
	return c.Props.Mode == DBInstaller
}

func (c *CreatorCtx) error(err error) *CreatorCtx {
	c.err = err
	return c
}

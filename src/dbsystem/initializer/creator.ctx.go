package initializer

import "HomegrownDB/dbsystem/config"

type CreatorCtx struct {
	Props CreatorProps

	RootPath string

	err error
}

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

func (c *CreatorCtx) initProperties() *CreatorCtx {

}

func (c *CreatorCtx) modeEqTest() bool {
	return c.Props.Mode == Test
}

func (c *CreatorCtx) modeEqDBInstaller() bool {
	return c.Props.Mode == DBInstaller
}

func (c *CreatorCtx) error(err error) *CreatorCtx {
	c.err = err
}

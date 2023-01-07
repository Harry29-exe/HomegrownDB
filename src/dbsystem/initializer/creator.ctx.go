package initializer

type CreatorCtx struct {
	Props CreatorProps

	RootPath string

	err error
}

func (c *CreatorCtx) initRootPath() *CreatorCtx {
	c.RootPath = c.Props.RootPath
	os.Set
}

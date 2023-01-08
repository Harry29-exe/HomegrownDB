package initializer

func InitializeDB(props CreatorProps) error {
	ctx := CreatorCtx{Props: props}
	ctx.initRootPath().
		initConfigurationAndProperties().
		initDBFilesystem().
		createRelations()
	return ctx.err
}

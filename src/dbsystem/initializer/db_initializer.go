package initializer

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/pageio"
	"os"
)

func InitializeDB(props *config.Properties) (dbsystem.DBSystem, error) {
	initFile, err := readInitProperties()
	if err != nil {
		return nil, err
	}
	pageIOStore := pageio.NewStore()

	initCtx := ctx{
		Props:       props,
		InitProps:   initFile,
		PageIOStore: pageIOStore,
	}
	_ = initCtx
	//todo implement me
	panic("Not implemented")
}

func initRelations(initCtx ctx) error {
	for _, rel := range initCtx.InitProps.Relations {
		fileData, err := readFile(initCtx.Props.DBHomePath + rel.RelativePath)
		if err != nil {
			return err
		}

		switch rel.RelKind {
		case relation.TypeTable:
			err = initTable(fileData, initCtx)
		}
	}
	//todo implement me
	panic("Not implemented")
}

func initTable(serializedData []byte, initCtx ctx) error {
	tableDef := table.Deserialize(serializedData)
	return initCtx.TableStore.LoadTable(tableDef)
}

//func initFSM(serializedData []byte, initCtx ctx) error {
//
//}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]byte, s.Size())
	_, err = f.Read(data)
	return data, err
}

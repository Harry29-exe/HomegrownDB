package dbfs

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/schema/table"
	"os"
)

func CreateTableFiles(table table.Definition) error {
	tablePath := dbsystem.DBHomePath() + "/tables/" + table.Name()
	err := os.Mkdir(tablePath, 755)
	if err != nil {
		return err
	}

	err = createDataFile(table, tablePath)
	if err != nil {
		return err
	}

	err = createBgDataFile(table, tablePath)
	if err != nil {
		return err
	}

	return nil
}

func serializeAndSave(table table.Definition, tablePath string) error {
	data := table.Serialize()

	file, err := os.Create(tablePath + "/info")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func createDataFile(table table.Definition, tablePath string) error {
	dataFile, err := os.Create(tablePath + "/data.hdbd")
	if err != nil {
		return err
	}
	err = dataFile.Close()
	if err != nil {
		return err
	}

	dataFSM, err := os.Create(tablePath + "/data_fsm.hdbfsm")
	if err != nil {
		return err
	}
	err = dataFSM.Close()
	if err != nil {
		return err
	}

	return nil
}

func createBgDataFile(table table.Definition, tablePath string) error {
	dataFile, err := os.Create(tablePath + "/bg_data.hdbd")
	if err != nil {
		return err
	}
	err = dataFile.Close()
	if err != nil {
		return err
	}

	dataFSM, err := os.Create(tablePath + "/bg_data_fsm.hdbfsm")
	if err != nil {
		return err
	}
	err = dataFSM.Close()
	if err != nil {
		return err
	}

	return nil
}

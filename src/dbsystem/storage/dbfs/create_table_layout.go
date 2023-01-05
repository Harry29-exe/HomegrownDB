package dbfs

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation/table"
	"os"
)

func serializeAndSave(table table.RDefinition, tablePath string) error {
	serialzier := bparse.NewSerializer()
	table.Serialize(serialzier)
	data := serialzier.GetBytes()

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

func createDataFile(table table.RDefinition, tablePath string) error {
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

func createBgDataFile(table table.RDefinition, tablePath string) error {
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

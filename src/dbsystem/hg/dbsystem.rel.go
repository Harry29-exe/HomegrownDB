package hg

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/pageio"
	"fmt"
)

var _ RelationsOperations = &DBSystem{}

func (db *DBSystem) CreateRel(rel relation.Relation) error {
	switch rel.Kind() {
	case relation.TypeTable:
		return db.createTable(rel.(table.Definition))
	case relation.TypeIndex:
		//todo implement me
		panic("Not implemented")
	default:
		panic(fmt.Sprintf("unknown relation type %+v", rel))
	}
}

func (db *DBSystem) createTable(tableDef table.Definition) (err error) {
	tableDef.SetOID(db.NextOID())
	tableDef.SetFsmOID(db.NextOID())

	if err = db.FS().InitNewRelationDir(tableDef.OID()); err != nil {
		return err
	}

	s := bparse.NewSerializer()
	tableDef.Serialize(s)
	if err = db.saveRelDefinition(tableDef.OID(), s.GetBytes()); err != nil {
		// todo db.FS.DeleteRelationDir(...)
		return err
	}

	if err = db.PageIOStore().Load(tableDef); err != nil {
		// todo db.FS.DeleteRelationDir(...)
		return err
	}

	if err = db.TableStore().LoadTable(tableDef); err != nil {
		// todo Page.Remove(tableDef)
		// todo db.FS.DeleteRelationDir(...)
		return err
	}

	if err = db.createFSM(tableDef.FsmOID()); err != nil {
		// todo Page.Remove(tableDef)
		// todo db.FS.DeleteRelationDir(...)
		// todo remove table from table store
		return err
	}
	return nil
}

func (db *DBSystem) createFSM(fsmOID dbobj.OID) (err error) {
	fs := db.FS()
	if err = fsm.CreateFreeSpaceMap(fsmOID, fs); err != nil {
		return err
	}

	file, err := fs.OpenRelationDataFile(fsmOID)
	if err != nil {
		//todo remove fsm
		return err
	}

	io, err := pageio.NewPageIO(file)
	if err != nil {
		//todo remove fsm & close file
		return err
	}
	db.PageIOStore().Register(fsmOID, io)
	db.FsmStore().Register(fsm.NewFSM(fsmOID, db.SharedBuffer()))

	return nil
}

func (db *DBSystem) saveRelDefinition(id relation.ID, definition []byte) (err error) {
	file, err := db.FS().OpenRelationDef(id)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = file.Close()
		} else {
			err = file.Close()
		}
	}()

	_, err = file.WriteAt(definition, 0)
	if err != nil {
		return err
	}
	err = db.FS().Truncate(file.Name(), int64(len(definition)))
	if err != nil {
		return err
	}
	return nil
}

func (db *DBSystem) LoadRel(rid relation.ID) error {
	data, err := db.readRelationDefFile(rid)
	if err != nil {
		return err
	}
	d := bparse.NewDeserializer(data)
	baseRel := relation.DeserializeBaseRelation(d)

	var rel relation.Relation
	switch baseRel.RelKind {
	case relation.TypeTable:
		return db.loadTable(data)
	case relation.TypeIndex:
		//todo implement me
		panic("Not implemented")
	default:
		panic(fmt.Sprintf("unknown relation type %+v", rel))
	}
}

func (db *DBSystem) loadTable(serializedTable []byte) error {
	tableDef := table.Deserialize(serializedTable)

	if err := db.PageIOStore().Load(tableDef); err != nil {
		return err
	}
	if err := db.TableStore().LoadTable(tableDef); err != nil {
		//todo delete table from pageio
		return err
	}
	return nil
}

func (db *DBSystem) readRelationDefFile(rid relation.ID) (data []byte, err error) {
	file, err := db.FS().OpenRelationDef(rid)
	defer func() {
		if err != nil {
			_ = file.Close()
		} else {
			err = file.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	data = make([]byte, stat.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return
}

func (db *DBSystem) DeleteRel(rel relation.Relation) error {
	//TODO implement me
	panic("implement me")
}

package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/access/table"
	"HomegrownDB/dbsystem/dbobj"
	"HomegrownDB/dbsystem/reldef"
	table2 "HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type Manager interface {
	Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error)
	Delete(relation reldef.Relation) error
	GetByOID(oid reldef.OID) reldef.Relation
	FindByName(name string) (reldef.OID, error)

	Lock(relationOID reldef.OID, mode LockMode)
	Unlock(relationOID reldef.OID, mode LockMode)
}

type LockMode uint8

const (
	LockRead LockMode = iota
	LockWrite
)

type OIDSequence interface {
	NextOID() reldef.OID
}

type tid struct {
	pageId  uint32
	tupleId uint16
}

type stdManager struct {
	Buffer      buffer.SharedBuffer
	FS          dbfs.FS
	OIDSequence OIDSequence

	oidMap  map[reldef.OID]tid
	nameMap map[string]reldef.OID
}

var _ Manager = &stdManager{}

func (s *stdManager) Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error) {
	if relation.OID() == dbobj.InvalidOID {
		relation.InitRel(s.OIDSequence.NextOID(), s.OIDSequence.NextOID(), s.OIDSequence.NextOID())
	}
	if err := s.createRelationOnDisc(relation); err != nil {
		return nil, err
	}

	switch relation.Kind() {
	case reldef.TypeTable:
		return relation, s.createTableInSysTables(relation.(table2.Definition), tx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (s *stdManager) createTableInSysTables(definition table2.Definition, tx tx.Tx) error {
	tuple := systable.RelationsOps.TableAsRelationsRow(definition, tx)
	err := table.Insert(tuple, tx, systable.RelationsTableDef(), fsm.NewFSM(systable.HGRelationsFsmOID, s.Buffer), s.Buffer)
	if err != nil {
		return err
	}
	columns := systable.ColumnsOps.DataToRows(definition.OID(), definition.Columns(), tx)
	for _, columnRow := range columns {
		err = table.Insert(columnRow, tx, systable.ColumnsTableDef(), fsm.NewFSM(systable.HGColumnsFsmOID, s.Buffer), s.Buffer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *stdManager) createRelationOnDisc(relation reldef.Relation) error {
	if err := s.FS.InitNewPageObjectDir(relation.OID()); err != nil {
		return err
	}
	if err := s.FS.InitNewPageObjectDir(relation.FsmOID()); err != nil {
		return err
	}
	if err := s.FS.InitNewPageObjectDir(relation.VmOID()); err != nil {
		return err
	}
	//todo add err handling (delete all created on fail)
	return nil
}

func (s *stdManager) Delete(relation reldef.Relation) error {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) GetByOID(oid reldef.OID) reldef.Relation {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) FindByName(name string) (reldef.OID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) Lock(relationOID reldef.OID, mode LockMode) {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) Unlock(relationOID reldef.OID, mode LockMode) {
	//TODO implement me
	panic("implement me")
}

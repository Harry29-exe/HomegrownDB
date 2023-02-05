package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/relation/dbobj"
	"HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

type Manager interface {
	Create(relation Relation, tx tx.Tx) (Relation, error)
	Delete(relation Relation) error
	FindByOID(oid OID) (Relation, error)
	FindByName(name string) (Relation, error)

	Lock(relationOID OID)
	Unlock(relationOID OID)
}

type OIDSequence interface {
	NextOID() OID
}

type tid struct {
	pageId  uint32
	tupleId uint16
}

type stdManager struct {
	Buffer      buffer.SharedBuffer
	FS          dbfs.FS
	OIDSequence OIDSequence

	oidMap  map[OID]tid
	nameMap map[string]tid
}

var _ Manager = &stdManager{}

func (s *stdManager) Create(relation Relation, tx tx.Tx) (Relation, error) {
	if relation.OID() == dbobj.InvalidOID {
		relation.InitRel(s.OIDSequence.NextOID(), s.OIDSequence.NextOID(), s.OIDSequence.NextOID())
	}
	if err := s.createRelationOnDisc(relation); err != nil {
		return nil, err
	}

	switch relation.Kind() {
	case TypeTable:
		return relation, s.createTableInSysTables(relation.(table.Definition), tx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (s *stdManager) createTableInSysTables(definition table.Definition, tx tx.Tx) error {
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

func (s *stdManager) createRelationOnDisc(relation Relation) error {
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

func (s *stdManager) Delete(relation Relation) error {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) FindByOID(oid OID) (Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) FindByName(name string) (Relation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) Lock(relationOID OID) {
	//TODO implement me
	panic("implement me")
}

func (s *stdManager) Unlock(relationOID OID) {
	//TODO implement me
	panic("implement me")
}

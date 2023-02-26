package relation

import (
	"HomegrownDB/dbsystem/access/sequence"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/access/table"
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/bparse"
	"log"
)

func (s *stdManager) Create(relation reldef.Relation, tx tx.Tx) (reldef.Relation, error) {
	s.mngLock.Lock()
	defer s.mngLock.Unlock()

	s.prepareRelation(relation)

	err := s.createRelationOnDisc(relation)
	if err != nil {
		return nil, err
	}

	switch relation.Kind() {
	case reldef.TypeTable:
		err = s.createTableInSysTables(relation.(reldef.TableDefinition), tx)
	case reldef.TypeSequence:
		err = s.createSequenceInSysTables(relation.(reldef.SequenceDef), tx)
		if err == nil {
			err = s.createSequenceFirstPage(relation.(reldef.SequenceDef))
		}
	default:
		//todo implement me
		panic("Not implemented")
	}
	if err != nil {
		return nil, err
	}

	s.cache.relations[relation.OID()] = relation
	s.nameMap[relation.Name()] = relation.OID()

	return relation, nil
}

func (s *stdManager) prepareRelation(relation reldef.Relation) {
	if relation.OID() != hglib.InvalidOID {
		return
	}

	relation.InitRel(s.nextOID(), s.nextOID(), s.nextOID())

	switch relation.Kind() {
	case reldef.TypeTable:
		s.prepareTableRel(relation)
	case reldef.TypeSequence, reldef.TypeIndex:
		// do nothing
	}
}

func (s *stdManager) prepareTableRel(relation reldef.Relation) {
	tableDef := relation.(reldef.TableDefinition)
	for _, col := range tableDef.Columns() {
		(col.(reldef.ColumnDefinition)).SetId(s.nextOID())
	}
}

func (s *stdManager) createTableInSysTables(definition reldef.TableDefinition, tx tx.Tx) error {
	tuple, err := systable.RelationsOps.TableAsRelationsRow(definition, tx)
	if err != nil {
		return err
	}
	err = table.Insert(tuple, tx, systable.RelationsTableDef(), fsm.NewFSM(systable.HGRelationsFsmOID, s.Buffer), s.Buffer)
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

func (s *stdManager) createSequenceInSysTables(definition reldef.SequenceDef, tx tx.Tx) error {
	relationsTuple, err := systable.RelationsOps.SequenceAsRelationsRow(definition, tx)
	if err != nil {
		return err
	}
	err = table.Insert(relationsTuple, tx, systable.RelationsTableDef(), fsm.NewFSM(systable.HGRelationsFsmOID, s.Buffer), s.Buffer)
	if err != nil {
		return err
	}

	sequencesTuple := systable.SequencesOps.DataToRow(definition, tx)
	err = table.Insert(sequencesTuple, tx, systable.SequencesTableDef(), fsm.NewFSM(systable.HGSequencesFsmOID, s.Buffer), s.Buffer)
	if err != nil {
		return err
	}

	return nil
}

func (s *stdManager) createSequenceFirstPage(definition reldef.SequenceDef) error {
	page, err := s.Buffer.SeqPage(definition.OID())
	if err != nil {
		return err
	}
	defer s.Buffer.WPageRelease(page.PageTag())

	switch definition.TypeTag() {
	case rawtype.TypeInt8:
		val := bparse.Serialize.Int8(definition.SeqStart())
		copy(page.Bytes, val)
	default:
		//todo implement me
		panic("Not implemented")
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

	if err := fsm.InitFreeSpaceMapFile(relation.FsmOID(), s.FS); err != nil {
		return err
	}

	//todo add err handling (delete all created on fail)
	return nil
}

func (s *stdManager) nextOID() reldef.OID {
	oidSequenceRel := s.cache.Get(systable.OIDSequenceOID)
	if oidSequenceRel == nil {
		log.Panicf("no sys sequence: sys_sequence_oid")
	}
	oidSequenceDef, ok := oidSequenceRel.(reldef.RSequenceDef)
	if !ok {
		log.Panicf("could not cast sys_sequence_oid sequence to reldef.RSequenceDef")
	}

	val, err := sequence.NextValue(oidSequenceDef, s.Buffer)
	if err != nil {
		log.Panicf("error occured: %s", err.Error())
	} else if val.TypeTag != rawtype.TypeInt8 {
		log.Panicf("sequence: sys_sequence_oid returned value other that int64")
	}

	return reldef.OID(bparse.Parse.Int8(val.NormValue))
}

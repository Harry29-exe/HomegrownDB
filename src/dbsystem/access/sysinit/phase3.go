package sysinit

import (
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/tx"
	"math"
)

var Phase3 = phase3{}

func (p3 phase3) Execute(manager relation.Manager) error {
	createTx := tx.StdTx{Id: 0}
	err := p3.createOIDSequence(manager, createTx)
	if err != nil {
		return err
	}

	return nil
}

func (phase3) createOIDSequence(mng relation.Manager, tx tx.Tx) error {
	oidSequence := reldef.NewSequenceDef(
		reldef.BaseRelation{
			RelName:          "sys_sequence_oid",
			ID:               systable.OIDSequenceOID,
			RelKind:          reldef.TypeSequence,
			FreeSpaceMapOID:  systable.OIDSequenceFsmOID,
			VisibilityMapOID: systable.OIDSequenceVmOID,
		},
		rawtype.TypeInt8,
		10_000,
		1,
		math.MaxInt64,
		0,
		0,
		false,
	)
	_, err := mng.Create(oidSequence, tx)

	return err
}

type phase3 struct{}

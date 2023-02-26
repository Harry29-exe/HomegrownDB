package sysinit

import (
	"HomegrownDB/dbsystem/access/relation"
)

var Phase2 = phase2{}

func (phase2) Execute(manager relation.Manager) error {
	//creatorTX := tx.StdTx{Id: 0}
	//_, err := manager.Create(systable.SequencesTableDef(), creatorTX)
	//if err != nil {
	//	return err
	//}

	return nil
}

type phase2 struct{}

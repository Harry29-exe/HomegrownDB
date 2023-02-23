package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var _ Builder = modifyTableBuilder{}

type modifyTableBuilder struct{}

func (m modifyTableBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	specificPlan := plan.(node.ModifyTable)
	resultRTE := ctx.GetRTE(specificPlan.ResultRelations[0])

	resultTable := resultRTE.Ref
	return &ModifyTable{
		Plan: specificPlan,
		Left: CreateFromPlan(specificPlan.Left, ctx),
		OutputPattern: page.TuplePattern{
			Columns: []page.PatternCol{
				{Type: hgtype.NewDefaultColType(rawtype.TypeInt8), Name: "Rows"},
			},
			BitmapLen: 1,
		},
		txCtx:       ctx.Tx,
		buff:        ctx.Buff,
		resultTable: resultTable,
		fsm:         fsm.NewFSM(resultTable.FsmOID(), ctx.Buff),
		done:        false,
	}
}

var _ ExecNode = &ModifyTable{}

type ModifyTable struct {
	Plan          node.ModifyTable
	Left          ExecNode
	OutputPattern page.TuplePattern

	txCtx       tx.Tx
	buff        buffer.SharedBuffer
	resultTable reldef.TableRDefinition
	fsm         *fsm.FSM
	done        bool
}

func (m *ModifyTable) Next() page.Tuple {
	tuplesInserted := int64(0)
	var err error

	for m.Left.HasNext() {
		tuple := m.Left.Next()
		err = m.tryInsert(tuple)
		for err != nil {
			err = m.tryInsert(tuple)
		}
		tuplesInserted++
	}

	m.done = true
	outputValues := [][]byte{intype.ConvInt8(tuplesInserted)}
	return page.NewTuple(outputValues, m.OutputPattern, m.txCtx)
}

func (m *ModifyTable) tryInsert(tuple page.Tuple) error {
	pageId, err := m.fsm.FindPage(uint16(tuple.TupleSize()), m.txCtx)
	if err != nil {
		log.Panic(err.Error())
	}
	wPage, err := m.buff.WTablePage(m.resultTable, pageId)
	if err != nil {
		panic(err.Error())
	}
	defer m.buff.WPageRelease(wPage.PageTag())

	return wPage.InsertTuple(tuple.Bytes())
}

func (m *ModifyTable) HasNext() bool {
	return !m.done
}

func (m *ModifyTable) Init(plan node.Plan) error {
	// ModifyTable does not need to be initialized
	return nil
}

func (m *ModifyTable) Close() error {
	// // ModifyTable does not need to be closed
	return nil
}

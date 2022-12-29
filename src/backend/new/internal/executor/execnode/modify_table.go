package execnode

import (
	"HomegrownDB/backend/new/internal/executor/exinfr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/inputtype"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/dpage"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
)

var _ Builder = modifyTableBuilder{}

type modifyTableBuilder struct{}

func (m modifyTableBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	specificPlan := plan.(node.ModifyTable)
	resultRTE := ctx.GetRTE(specificPlan.ResultRelations[0])

	return &ModifyTable{
		Plan: specificPlan,
		Left: CreateFromPlan(specificPlan.Left, ctx),
		OutputPattern: &dpage.TuplePattern{
			Columns:   []dpage.ColumnInfo{{CType: hgtype.NewInt8(hgtype.Args{})}},
			BitmapLen: 1,
		},
		txCtx:       ctx.TxCtx,
		buff:        ctx.Buff,
		resultTable: resultRTE.Ref,
		fsm:         ctx.FsmStore.GetFsmFor(resultRTE.TableId),
		done:        false,
	}
}

var _ ExecNode = &ModifyTable{}

type ModifyTable struct {
	Plan          node.ModifyTable
	Left          ExecNode
	OutputPattern *dpage.TuplePattern

	txCtx       *tx.Ctx
	buff        buffer.SharedBuffer
	resultTable table.RDefinition
	fsm         *fsm.FreeSpaceMap
	done        bool
}

func (m *ModifyTable) Next() dpage.Tuple {
	tuplesInserted := int64(0)
	for m.Left.HasNext() {
		tuple := m.Left.Next()
		for {
			pageId, err := m.fsm.FindPage(uint16(tuple.TupleSize()), m.txCtx)
			if err != nil {
				panic(err.Error())
			}
			wPage, err := m.buff.WTablePage(m.resultTable, pageId)
			if err != nil {
				panic(err.Error())
			}

			err = wPage.InsertTuple(tuple.Data())
			if err == nil {
				break
			}
		}
		tuplesInserted++
	}

	outputValues := [][]byte{inputtype.ConvInt8(tuplesInserted)}
	return dpage.NewTuple(outputValues, m.OutputPattern, m.txCtx)
}

func (m *ModifyTable) HasNext() bool {
	return !m.done
}

func (m *ModifyTable) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}

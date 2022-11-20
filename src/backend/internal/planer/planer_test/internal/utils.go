package internal

import (
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/common/tests/assert"
	"testing"
)

type ANode struct {
	T *testing.T
}

//func (n ANode) CmpSelectField(pnode pnode.FieldNode, anode anode.SelectField) {
//	assert.Eq(pnode.FieldName, anode.Column.Name(), n.T)TablesNode
//	assert.Eq(pnode.FieldAlias, anode.FieldAlias, n.T)
//	assert.Eq(pnode.TableAlias, anode.)
//}

func (n ANode) CmpTables(pnodeTables []pnode.TableNode, anodeTables []qctx.QTableId, ctx qctx.QueryCtx) {
	assert.Eq(len(pnodeTables), len(anodeTables), n.T)
	for i := 0; i < len(pnodeTables); i++ {
		pnodeTable := pnodeTables[i]

		//assert alias
		qTableId := ctx.QTCtx.GetQTableId(pnodeTable.TableAlias)
		assert.NotEq(qTableId, qctx.InvalidQTableId, n.T)
		//assert table
		anodeTable := ctx.QTCtx.GetTableByQTableId(anodeTables[i])
		assert.Eq(pnodeTable.TableName, anodeTable.Name(), n.T)
	}
}

func (n ANode) ValidateFields(pnodeFields []pnode.FieldNode, anodeFields []qctx.QColumnId, ctx qctx.QueryCtx) {
	assert.Eq(len(pnodeFields), len(anodeFields), n.T)
	for i := 0; i < len(pnodeFields); i++ {
		pnFields := pnodeFields[i]
		anField := anodeFields[i]
		table := ctx.QTCtx.GetTableByQTableId(anField.QTableId)
		assert.Eq(pnFields.FieldAlias)
	}
}

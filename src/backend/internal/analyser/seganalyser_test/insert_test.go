package seganalyser

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/relation/table"
	"testing"
)

func TestInsertAnalyse_SimplePositive1(t *testing.T) {
	//given
	query := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	store, users := TestTableStore.StoreWithUsersTable(t)
	expectedNode := InsertTests.expectedSimplePositive1(users, t)

	//when
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	//then
	NodeAssert.Eq(expectedNode, queryNode, t)
}

func (insertTest) expectedSimplePositive1(users table.Definition, t *testing.T) node2.Query {
	rteIdCounter := appsync.NewSimpleCounter[node2.RteID](0)

	query := node2.NewQuery(node2.CommandTypeInsert, nil)

	resultRel := node2.NewRelationRTE(rteIdCounter.Next(), users)
	query.RTables = []node2.RangeTableEntry{resultRel}
	query.ResultRel = resultRel.Id

	valuesRte := node2.NewValuesRTE(rteIdCounter.Next(), [][]node2.Expr{
		{node2.NewConstInt8(1), NewConstStr("bob", t)},
	})
	valuesRte.ColTypes = []coltype.ColumnType{
		coltype.NewInt8(hgtype.Args{}),
		coltype.NewStr(hgtype.Args{
			Length: uint32(len("bob")),
			VarLen: true,
			UTF8:   false,
		}),
	}
	query.TargetList = []node2.TargetEntry{
		node2.NewTargetEntry(node2.NewVar(valuesRte.Id, 0, users.Column(0).CType()), tt_user.C0IdOrder, tt_user.C0Id),
		node2.NewTargetEntry(node2.NewConst(tt_user.C1AgeType.Tag, nil), tt_user.C1AgeOrder, tt_user.C1Age),
		node2.NewTargetEntry(node2.NewVar(valuesRte.Id, 1, users.Column(2).CType()), tt_user.C2NameOrder, tt_user.C2Name),
		node2.NewTargetEntry(node2.NewConst(tt_user.C3SurnameType.Tag, nil), tt_user.C3SurnameOrder, tt_user.C3Surname),
	}

	query.AppendRTE(valuesRte)
	query.FromExpr = node2.NewFromExpr2(nil, valuesRte.CreateRef())

	return query
}

var InsertTests = insertTest{}

type insertTest struct{}

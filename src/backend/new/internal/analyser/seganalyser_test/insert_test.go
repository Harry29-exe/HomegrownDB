package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/table"
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

func (insertTest) expectedSimplePositive1(users table.Definition, t *testing.T) node.Query {
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	query := node.NewQuery(node.CommandTypeInsert, nil)

	resultRel := node.NewRelationRTE(rteIdCounter.Next(), users)
	query.RTables = []node.RangeTableEntry{resultRel}
	query.ResultRel = resultRel.Id

	valuesRte := node.NewValuesRTE(rteIdCounter.Next(), [][]node.Expr{
		{node.NewConstInt8(1), NewConstStr("bob", t)},
	})
	valuesRte.ColTypes = []hgtype.TypeData{
		hgtype.NewInt8(hgtype.Args{}),
		hgtype.NewStr(hgtype.Args{
			Length: uint32(len("bob")),
			VarLen: true,
			UTF8:   false,
		}),
	}
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 0, users.Column(0).CType()), tt_user.C0IdOrder, tt_user.C0Id),
		node.NewTargetEntry(node.NewConst(tt_user.C1AgeType.Tag, nil), tt_user.C1AgeOrder, tt_user.C1Age),
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 1, users.Column(2).CType()), tt_user.C2NameOrder, tt_user.C2Name),
		node.NewTargetEntry(node.NewConst(tt_user.C3SurnameType.Tag, nil), tt_user.C3SurnameOrder, tt_user.C3Surname),
	}

	query.AppendRTE(valuesRte)
	query.FromExpr = node.NewFromExpr2(nil, valuesRte.CreateRef())

	return query
}

var InsertTests = insertTest{}

type insertTest struct{}

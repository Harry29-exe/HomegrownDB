package analyse_test

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/datastructs/appsync"
	"HomegrownDB/lib/tests/assert"
	"HomegrownDB/lib/tests/tutils/testtable/tt_user"
	"testing"
)

func TestInsertAnalyse_SimplePositive1(t *testing.T) {
	//given
	query := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	db := hgtest.CreateAndLoadDBWith(nil, t).
		WithUsersTable().
		Build()

	store := db.DB.AccessModule().RelationManager()
	users := db.TableByName(tt_user.TableName)
	expectedNode := InsertTests.expectedSimplePositive1(users, t)

	//when
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	//then
	NodeAssert.Eq(expectedNode, queryNode, t)
}

func (insertTest) expectedSimplePositive1(users tabdef.TableRDefinition, t *testing.T) node.Query {
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	query := node.NewQuery(node.CommandTypeInsert, nil)

	resultRel := node.NewRelationRTE(rteIdCounter.Next(), users)
	query.RTables = []node.RangeTableEntry{resultRel}
	query.ResultRel = resultRel.Id

	valuesRte := node.NewValuesRTE(rteIdCounter.Next(), [][]node.Expr{
		{node.NewConstInt8(1), NewConstStr("bob", t)},
	})
	valuesRte.ColTypes = []hgtype.ColumnType{
		hgtype.NewInt8(rawtype.Args{}),
		hgtype.NewStr(rawtype.Args{
			Length: len("bob"),
			VarLen: true,
			UTF8:   false,
		}),
	}
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 0, users.Column(0).CType()), tt_user.C0IdOrder, tt_user.C0Id),
		node.NewTargetEntry(node.NewConst(tt_user.C1AgeType.ColTag, nil), tt_user.C1AgeOrder, tt_user.C1Age),
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 1, users.Column(2).CType()), tt_user.C2NameOrder, tt_user.C2Name),
		node.NewTargetEntry(node.NewConst(tt_user.C3SurnameType.ColTag, nil), tt_user.C3SurnameOrder, tt_user.C3Surname),
	}

	query.AppendRTE(valuesRte)
	query.FromExpr = node.NewFromExpr2(nil, valuesRte.CreateRef())

	return query
}

var InsertTests = insertTest{}

type insertTest struct{}

package node

import (
	"HomegrownDB/backend/internal/pnode"
	tabdef "HomegrownDB/dbsystem/reldef/tabdef"
	"fmt"
	"log"
	"reflect"
)

// -------------------------
//      CommandType
// -------------------------

type CommandType uint8

const (
	// CommandTypeSelect for select statements
	CommandTypeSelect CommandType = iota
	// CommandTypeInsert for insert statements
	CommandTypeInsert
	// CommandTypeUpdate for update statements
	CommandTypeUpdate
	// CommandTypeDelete for delete statements
	CommandTypeDelete
	// CommandTypeUtils for statements that operate
	// rater on db structure than on db data
	CommandTypeUtils
)

func (c CommandType) ToString() string {
	return [...]string{
		"CommandTypeSelect",
		"CommandTypeUpdate",
		"CommandTypeDelete",
		"CommandTypeUtils",
	}[c]
}

// -------------------------
//      Query
// -------------------------

func NewQuery(commandType CommandType, srcStmt pnode.Node) Query {
	return &query{
		node:    node{tag: TagQuery},
		Command: commandType,
		SrcStmt: srcStmt,
	}
}

var _ Node = &query{}

type Query = *query

type query struct {
	node
	Command   CommandType
	SrcStmt   pnode.Node // stmt that was used to create this Query
	UtilsStmt Node       // UtilsStmt not nil only when Command == CommandTypeUtils

	TargetList []TargetEntry

	ResultRel RteID             // Id of result tabdef, for insert, update, delete
	RTables   []RangeTableEntry // Tables used in Query
	FromExpr  FromExpr
}

func (q Query) dEqual(node Node) bool {

	raw := node.(Query)
	return q.Command == raw.Command &&
		DEqual(q.UtilsStmt, raw.UtilsStmt) &&
		cmpNodeArray(q.TargetList, raw.TargetList) &&
		q.ResultRel == raw.ResultRel &&
		cmpNodeArray(q.RTables, raw.RTables) &&
		DEqual(q.FromExpr, raw.FromExpr)
}

func (q query) DPrint(nesting int) string {
	n1 := nesting + 1
	return fmt.Sprintf(
		`%s{
Command: 	%s, 
TargetList: %s,
ResultRel: 	%d,
RTables: 	%s
FromExpr: 	%s`,
		q.dTag(nesting),
		q.Command.ToString(),
		dPrintArr(n1, q.TargetList),
		q.ResultRel,
		dPrintArr(n1, q.RTables),
		q.FromExpr.DPrint(n1),
	)
}

func (q Query) GetRTE(id RteID) RangeTableEntry {
	for i := 0; i < len(q.RTables); i++ {
		if q.RTables[i].Id == id {
			return q.RTables[i]
		}
	}
	return nil
}

func (q Query) AppendRTE(rte RangeTableEntry) {
	q.RTables = append(q.RTables, rte)
}

// -------------------------
//      Commands
// -------------------------

func NewCreateRelationTable(table tabdef.Definition) CreateRelation {
	return &createRelation{
		node: node{
			tag: TagCreateTable,
		},
		FutureTable: table,
	}
}

type CreateRelation = *createRelation

var _ Node = &createRelation{}

type createRelation struct {
	node
	FutureTable tabdef.Definition
	FutureIndex any // not supported yet
}

func (c CreateRelation) dEqual(node Node) bool {
	raw := node.(CreateRelation)

	return c.tablesEq(raw)
}

func (c CreateRelation) tablesEq(raw CreateRelation) bool {
	t1, t2 := c.FutureTable, raw.FutureTable
	if len(t1.Columns()) != len(t2.Columns()) {
		log.Printf("CreateTable.FutureTable were not equal: columns length are different")
		return false
	}
	for colNo := 0; colNo < len(t1.Columns()); colNo++ {
		c1, c2 := t1.Columns()[colNo], t2.Columns()[colNo]
		if !reflect.DeepEqual(c1, c2) {
			log.Printf("CreateTable.FutureTable were not equal: \nexpected col: %+v\nactual col: %+v",
				c1, c2)
			return false
		}
	}
	return reflect.DeepEqual(t1, t2)
}

func (c CreateRelation) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

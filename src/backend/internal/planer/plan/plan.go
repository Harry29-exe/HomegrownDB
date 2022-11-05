package plan

import (
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/dbsystem/schema/table"
)

func NewPlan() Plan {
	return &plan{
		rootNode: nil,
		tables:   make([]Table, 10),
	}
}

type Plan interface {
	RootNode() Node
	SetRootNode(node Node)

	AddTable(analyserTable anode.Table) Table
	Tables() []Table
}

type plan struct {
	rootNode Node

	tables []Table
}

func (p *plan) RootNode() Node {
	return p.rootNode
}

func (p *plan) SetRootNode(node Node) {
	p.rootNode = node
}

func (p *plan) AddTable(analyserTable anode.Table) Table {
	tab := Table{
		TableId:     analyserTable.Def.TableId(),
		PlanTableId: TableId(analyserTable.QTableId),
	}
	p.tables = append(p.tables, tab)
	return tab
}

func (p *plan) Tables() []Table {
	return p.tables
}

type Table struct {
	TableId     table.Id
	PlanTableId TableId
}

type TableId = uint8

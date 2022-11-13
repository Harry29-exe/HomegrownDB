package plan

import (
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

func (p *plan) Tables() []Table {
	return p.tables
}

type Table struct {
	TableId   table.Id
	TableHash string
}

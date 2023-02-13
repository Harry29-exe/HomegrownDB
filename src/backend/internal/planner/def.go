package planner

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/lib/datastructs/appsync"
)

type PlanNodeIcCounter = appsync.SimpleSyncCounter[node2.PlanNodeId]

func Plan(query node2.Query) (node2.PlanedStmt, error) {
	planedStmt := node2.NewPlanedStmt(query.Command)
	rootState := NewRootState(planedStmt)

	planTree, err := delegate(query, rootState)

	if err != nil {
		return planedStmt, err
	}
	planedStmt.PlanTree = planTree
	return planedStmt, nil
}

func delegate(query node2.Query, parentState State) (node2.Plan, error) {
	switch query.Command {
	case node2.CommandTypeSelect:
		return Select.Plan(query, parentState)
	case node2.CommandTypeInsert:
		return Insert.Plan(query, parentState)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

type PlanNodeCounter = appsync.SimpleSyncCounter[node2.PlanNodeId]

func NewRootState(planedStmt node2.PlanedStmt) State {
	return State{
		Root:            planedStmt,
		PlanNodeCounter: appsync.NewSimpleCounter[node2.PlanNodeId](0),
		ParentState:     nil,
		Plan:            nil,
		Query:           nil,
	}
}

type State struct {
	Root            node2.PlanedStmt
	PlanNodeCounter PlanNodeCounter
	ParentState     *State

	Plan  node2.Plan
	Query node2.Query
}

func (s State) CreateChildState(query node2.Query, plan node2.Plan) State {
	return State{
		Root:            s.Root,
		PlanNodeCounter: s.PlanNodeCounter,
		ParentState:     &s,

		Plan:  plan,
		Query: query,
	}
}

func (s State) NextPlanNodeId() node2.PlanNodeId {
	return s.PlanNodeCounter.Next()
}

func (s State) AppendRTE(rte ...node2.RangeTableEntry) {
	s.Root.Tables = append(s.Root.Tables, rte...)
}

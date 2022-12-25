package planner

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/common/datastructs/appsync"
)

type PlanNodeIcCounter = appsync.SimpleSyncCounter[node.PlanNodeId]

func Plan(query node.Query) (node.PlanedStmt, error) {
	planedStmt := node.NewPlanedStmt(query.Command)
	rootState := NewRootState(planedStmt)

	planTree, err := delegate(query, rootState)

	if err != nil {
		return planedStmt, err
	}
	planedStmt.PlanTree = planTree
	return planedStmt, nil
}

func delegate(query node.Query, parentState State) (node.Plan, error) {
	switch query.Command {
	case node.CommandTypeSelect:
		return Select.Plan(query, parentState)
	case node.CommandTypeInsert:
		return Insert.Plan(query, parentState)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

type PlanNodeCounter = appsync.SimpleSyncCounter[node.PlanNodeId]

func NewRootState(planedStmt node.PlanedStmt) State {
	return State{
		Root:            planedStmt,
		PlanNodeCounter: appsync.NewSimpleCounter[node.PlanNodeId](0),
		ParentState:     nil,
		Plan:            nil,
		Query:           nil,
	}
}

type State struct {
	Root            node.PlanedStmt
	PlanNodeCounter PlanNodeCounter
	ParentState     *State

	Plan  node.Plan
	Query node.Query
}

func (s State) CreateChildState(query node.Query, plan node.Plan) State {
	return State{
		Root:            s.Root,
		PlanNodeCounter: s.PlanNodeCounter,
		ParentState:     &s,

		Plan:  plan,
		Query: query,
	}
}

func (s State) NextPlanNodeId() node.PlanNodeId {
	return s.PlanNodeCounter.Next()
}

func (s State) AppendRTE(rte ...node.RangeTableEntry) {
	s.Root.Tables = append(s.Root.Tables, rte...)
}

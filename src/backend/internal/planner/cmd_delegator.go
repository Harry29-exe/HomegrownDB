package planner

import (
	"HomegrownDB/backend/internal/node"
	"log"
)

var CmdDelegator = cmdDelegator{}

type cmdDelegator struct{}

func (cmdDelegator) Plan(query node.Query, parentState State) (node.Plan, error) {
	stmt := query.UtilsStmt
	if stmt == nil {
		log.Panicf("utils statement is nil")
	}

	switch stmt.Tag() {
	case node.TagCreateRelation:
		return CreateRelation.Plan(stmt.(node.CreateRelation), parentState)
	default:
		log.Panicf("not supported command: %s", stmt.Tag().ToString())
	}
	return nil, nil
}

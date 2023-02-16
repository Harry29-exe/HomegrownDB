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

	//todo implement me
	panic("Not implemented")
	//switch stmt.Tag() {
	//case node.TagCreateTable:
	//	return CreateRelation.Plan(stmt.(node.CreateTable))
	//}
}

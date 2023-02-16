package node

import "HomegrownDB/dbsystem/reldef"

type CreateRelationPlan = *createRelationPlan

var _ Node = &createRelationPlan{}

type createRelationPlan struct {
	node
	RelationToCreate reldef.Relation
}

func (c *createRelationPlan) dEqual(node Node) bool {
	//TODO implement me
	panic("implement me")
}

func (c *createRelationPlan) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

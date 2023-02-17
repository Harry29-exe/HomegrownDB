package node

import (
	"HomegrownDB/dbsystem/reldef"
	"reflect"
)

// -------------------------
//      CreateRelationPlan
// -------------------------

func NewCreateRelationPlan(relationToCreate reldef.Relation, planNodeId PlanNodeId, querySrc Query) CreateRelationPlan {
	return &createRelationPlan{
		plan:             newPlan(TagCreateRelationPlan, planNodeId, querySrc),
		RelationToCreate: relationToCreate,
	}
}

type CreateRelationPlan = *createRelationPlan

var _ Plan = &createRelationPlan{}

type createRelationPlan struct {
	plan
	RelationToCreate reldef.Relation
}

func (c CreateRelationPlan) dEqual(node Node) bool {
	raw := node.(CreateRelationPlan)
	return reflect.DeepEqual(
		c.RelationToCreate, raw.RelationToCreate)
}

func (c CreateRelationPlan) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

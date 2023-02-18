package execnode

import (
	"HomegrownDB/backend/internal/executor/exinfr"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var _ Builder = createRelationBuilder{}

type createRelationBuilder struct{}

func (c createRelationBuilder) Create(plan node.Plan, ctx exinfr.ExCtx) ExecNode {
	return &CreateRelation{
		Plan: plan.(node.CreateRelationPlan),
		OutputPattern: page.NewPattern([]page.PatternCol{
			{Type: hgtype.NewInt8(hgtype.Args{Nullable: false}), Name: "rowsAffected"},
		}),
		Tx:         ctx.Tx,
		RelManager: ctx.TableStore,
		Done:       false,
	}
}

type CreateRelation struct {
	Plan          node.CreateRelationPlan
	OutputPattern page.TuplePattern

	Tx         tx.Tx
	RelManager relation.Manager

	Done bool
}

var _ ExecNode = &CreateRelation{}

func (c *CreateRelation) Next() page.Tuple {
	_, err := c.RelManager.Create(c.Plan.RelationToCreate, c.Tx)
	if err != nil {
		log.Panicf("error occured: %s", err.Error())
	}

	c.Done = true
	return page.NewTuple([][]byte{intype.ConvInt8(1)}, c.OutputPattern, c.Tx)
}

func (c *CreateRelation) HasNext() bool {
	return !c.Done
}

func (c *CreateRelation) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}

func (c *CreateRelation) Shutdown() error {
	//TODO implement me
	panic("implement me")
}

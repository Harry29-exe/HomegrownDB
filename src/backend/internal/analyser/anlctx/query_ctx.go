package anlctx

import (
	"HomegrownDB/backend/internal/node"
)

type QueryCtx = *queryCtx

func NewQueryCtx(query node.Query, ctx Ctx) QueryCtx {
	return &queryCtx{
		Ctx:       ctx,
		Query:     query,
		ParentCtx: nil,
	}
}

type queryCtx struct {
	Ctx
	Query     node.Query
	ParentCtx *queryCtx
}

func (q QueryCtx) IsRoot() bool {
	return q.ParentCtx == nil
}

func (q QueryCtx) CreateChildCtx(query node.Query) QueryCtx {
	return &queryCtx{
		Ctx:       q.Ctx,
		Query:     query,
		ParentCtx: q,
	}
}

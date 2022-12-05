package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var From = from{}

type from struct{}

func (f from) Analyse(query node.Query, stmt pnode.Node, ctx query.Ctx) error {
	switch query.Command {
	case node.CommandTypeSelect:
		return FromSelect
	}
}

// -------------------------
//      FromSelect
// -------------------------

var FromSelect = fromSelect{}

type fromSelect struct{}

func (f fromSelect) Analyse(query node.Query, stmt pnode.SelectStmt, ctx query.Ctx) error {
	stmt.From
}

// -------------------------
//      FromInsert
// -------------------------

// -------------------------
//      FromUtils
// -------------------------

var FromUtils = fromUtils{}

type fromUtils struct{}

package analyser

import (
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/stores"
	"HomegrownDB/dbsystem/tx"
	"fmt"
)

func Analyse(tree parser.Tree, txCtx tx.Ctx, tableStore stores.RTablesDefs) (Tree, error) {
	ctx := internal.NewAnalyserCtx(txCtx, tableStore)
	root, rootNodeType, err := delegateAnalyse(tree, ctx)
	if err != nil {
		return Tree{}, err
	}

	return Tree{
		RootType: rootNodeType,
		Root:     root,
	}, nil
}

func delegateAnalyse(tree parser.Tree, ctx *internal.AnalyserCtx) (root any, rootType rootType, err error) {
	switch tree.RootType {
	case parser.Select:
		rootType = RootTypeSelect
		rootPnode, ok := tree.Root.(pnode.Select)
		if !ok {
			panic(fmt.Sprintf("Could not cast pnode with root type of %d", tree.RootType))
		}
		root, err = seganalyser.Select.Analyse(rootPnode, ctx)
		return
	case parser.Insert:
		rootType = RootTypeInsert
		//todo implement me
		panic("Not implemented")
	default:
		//todo implement me
		panic("Not implemented")
	}
}

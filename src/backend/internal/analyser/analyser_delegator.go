package analyser

import (
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"fmt"
)

func Analyse(tree parser.Tree, ctx qctx.QueryCtx) (Tree, error) {
	root, rootNodeType, err := delegateAnalyse(tree, ctx)
	if err != nil {
		return Tree{}, err
	}

	return Tree{
		RootType: rootNodeType,
		Root:     root,
	}, nil
}

func delegateAnalyse(tree parser.Tree, ctx qctx.QueryCtx) (root any, rootType rootType, err error) {
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

package analyser

import (
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/dbsystem/stores"
	"fmt"
)

func Analyse(tree parser.Tree) (Tree, error) {
	root, rootNodeType, err := delegateAnalyse(tree, stores.DBTables)
	if err != nil {
		return Tree{}, err
	}

	return Tree{
		RootType: rootNodeType,
		Root:     root,
	}, nil
}

func delegateAnalyse(tree parser.Tree, store *stores.TablesStore) (root any, rootType rootType, err error) {
	switch tree.RootType {
	case parser.Select:
		rootType = RootTypeSelect
		rootPnode, ok := tree.Root.(pnode.SelectNode)
		if !ok {
			panic(fmt.Sprintf("Could not cast pnode with root type of %d", tree.RootType))
		}
		root, err = internal.Select.Analyse(rootPnode, store)
		return
	default:
		//todo implement me
		panic("Not implemented")
	}
}

package analyser

import (
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

type Tree struct {
	RootType rootType
	Root     any
}

type rootType = uint8

const (
	Select rootType = iota
	Insert
)

func delegateAnalyse(tree parser.Tree, store *stores.TablesStore) (root any, rootType rootType, err error) {
	switch tree.RootType {
	case parser.Select:
		rootType = Select
		rootPnode, ok := tree.Root.(pnode.SelectNode)
		if !ok {
			panic(fmt.Sprintf("Could not cast pnode with root type of %d", tree.RootType))
		}
		root, err = AnalyseSelect(rootPnode, store)
		return
	default:
		//todo implement me
		panic("Not implemented")
	}
}

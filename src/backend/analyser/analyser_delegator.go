package analyser

import (
	"HomegrownDB/backend/parser"
	"HomegrownDB/backend/parser/pnode"
	"HomegrownDB/dbsystem/stores"
	"fmt"
)

func Analyse(tree parser.Tree) (Tree, error) {
	root, rootType, err := delegateAnalyse(tree, stores.DBTables)
	if err != nil {
		return Tree{}, err
	}

	return Tree{
		RootType: rootType,
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
	switch tree.Type() {
	case parser.SELECT:
		rootType = Select
		rootPnode, ok := tree.Root().(pnode.SelectNode)
		if !ok {
			panic(fmt.Sprintf("Could not cast pnode with root type of %d", tree.Type()))
		}
		root, err = AnalyseSelect(rootPnode, store)
		return
	default:
		//todo implement me
		panic("Not implemented")
	}
}

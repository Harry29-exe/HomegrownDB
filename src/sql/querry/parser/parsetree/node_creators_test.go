package parsetree_test

import (
	"HomegrownDB/sql/querry/parser/parsetree"
	"testing"
)

func TestNewSelectNode(t *testing.T) {
	node := parsetree.NewSelectNode()

	illegalToAddNode := parsetree.NewSelectNode()
	err := node.AddChild(illegalToAddNode)
	if err == nil {
		t.Error("select node did returned error when adding illegal node")
		t.FailNow()
	} else if len(node.Children()) != 0 {
		t.Error("select node added illegal node to it's list without error")
		t.FailNow()
	}

	fieldsNode := parsetree.NewFieldsNode()
	if fieldsNode.Type() != parsetree.Fields {
		t.FailNow()
	}
	err = node.AddChild(fieldsNode)
	if err != nil {
		t.Error("select node did not accept field node")
	}
}

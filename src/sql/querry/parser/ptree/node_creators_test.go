package ptree_test

import (
	"HomegrownDB/sql/querry/parser/ptree"
	"testing"
)

func TestNewSelectNode(t *testing.T) {
	node := ptree.NewSelectNode()

	illegalToAddNode := ptree.NewSelectNode()
	err := node.AddChild(illegalToAddNode)
	if err == nil {
		t.Error("select node did returned error when adding illegal node")
		t.FailNow()
	} else if len(node.Children()) != 0 {
		t.Error("select node added illegal node to it's list without error")
		t.FailNow()
	}

	fieldsNode := ptree.NewFieldsNode()
	if fieldsNode.Type() != ptree.Fields {
		t.FailNow()
	}
	err = node.AddChild(fieldsNode)
	if err != nil {
		t.Error("select node did not accept field node")
	}
}

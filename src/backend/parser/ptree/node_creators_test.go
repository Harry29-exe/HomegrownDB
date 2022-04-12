package ptree_test

import (
	ptree2 "HomegrownDB/backend/parser/ptree"
	"testing"
)

func TestNewSelectNode(t *testing.T) {
	node := ptree2.NewSelectNode()

	illegalToAddNode := ptree2.NewSelectNode()
	err := node.AddChild(illegalToAddNode)
	if err == nil {
		t.Error("select node did returned error when adding illegal node")
		t.FailNow()
	} else if len(node.Children()) != 0 {
		t.Error("select node added illegal node to it's list without error")
		t.FailNow()
	}

	fieldsNode := ptree2.NewFieldsNode()
	if fieldsNode.Type() != ptree2.Fields {
		t.FailNow()
	}
	err = node.AddChild(fieldsNode)
	if err != nil {
		t.Error("select node did not accept field node")
	}
}

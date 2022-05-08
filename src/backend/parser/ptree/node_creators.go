package ptree

func NewSelectNode() Node {
	return &basicNode{
		acceptingNodeTypes: acceptedBySelectNode,
		children:           make([]Node, 0, 8),
		value:              nil,
		nodeType:           Select,
	}
}

var acceptedBySelectNode = []NodeType{Fields, From, Where}

func NewFromNode() Node {
	return &basicNode{
		acceptingNodeTypes: acceptedByFromNode,
		children:           make([]Node, 0, 2),
		value:              nil,
		nodeType:           Table,
	}
}

var acceptedByFromNode = []NodeType{Table}

func NewFieldsNode() Node {
	return &basicNode{
		acceptingNodeTypes: acceptedByFieldsNode,
		children:           make([]Node, 0, 10),
		value:              nil,
		nodeType:           Fields,
	}
}

var acceptedByFieldsNode = []NodeType{Field}

func NewFieldNode(value FieldNodeValue) Node {
	return &leafNode{
		value:    value,
		nodeType: Field,
	}
}

func NewTableNode(value TableNodeValue) Node {
	return &leafNode{
		value:    value,
		nodeType: Table,
	}
}

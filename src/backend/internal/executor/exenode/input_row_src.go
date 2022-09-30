package exenode

type InputRowSrc struct {
	Fields []inputRowField
}

type inputRowField struct {
	Values     []byte
	Expression ExeNode
}

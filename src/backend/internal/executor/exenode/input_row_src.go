package exenode

type InputRowSrc struct {
	Fields []inputRowField
}

func (rs *InputRowSrc) NextRow() [][]byte {
	row := make([][]byte, len(rs.Fields))
	for i, field := range rs.Fields {
		row[i] = field.NextValue()
	}

	return row
}

type inputRowField struct {
	Values     [][]byte
	Expression ExeNode
}

func (f *inputRowField) NextValue() (val []byte) {
	if f.Values != nil {
		val = f.Values[0]
		if len(f.Values) == 1 {
			f.Values = nil
		} else {
			f.Values = f.Values[1:]
		}
		return
	}

	qrow := f.Expression.Next()

}

package handler

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/storage/page"
	"bytes"
	"fmt"
	"log"
)

func ToJsonResult(tuples []page.RTuple) SqlResult {
	if len(tuples) == 0 {
		return JsonResult{data: []byte("[]")}
	}
	json := resultJsonSerializer{}

	json.WriteRune('[')
	pattern := tuples[0].Pattern()
	for _, tuple := range tuples {
		json.serializeRow(tuple, pattern.Columns)
	}
	json.WriteRune(']')
	return JsonResult{data: json.Bytes()}
}

type resultJsonSerializer struct {
	bytes.Buffer
	rowsWritten int
}

func (s *resultJsonSerializer) serializeRow(tuple page.RTuple, columns []page.PatternCol) {
	if s.rowsWritten > 0 {
		s.Buffer.WriteRune(',')
	}
	s.Buffer.WriteRune('{')
	s.Buffer.WriteString(fmt.Sprintf("\"%s\":", columns[0].Name))
	s.serializeValue(tuple.ColValue(0), columns[0].Type)

	for _, column := range columns[1:] {
		s.Buffer.WriteRune(',')
		s.Buffer.WriteString(fmt.Sprintf("\"%s\":", column.Name))
	}
	s.Buffer.WriteRune('}')
}

func (s *resultJsonSerializer) serializeValue(value []byte, t hgtype.Type) {
	switch t.Tag() {
	case hgtype.TypeStr:
		s.Buffer.WriteString(fmt.Sprintf("\"%s\"", string(value)))
	case hgtype.TypeInt8:
		s.Buffer.WriteString(fmt.Sprintf("%d", bparse.Parse.Int8(value)))
	default:
		log.Panic("unsupported hgtype type: " + t.Tag().ToStr())
	}
}

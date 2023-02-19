package handler

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/lib/bparse"
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
	for tupleNo, tuple := range tuples {
		if tupleNo != 0 {
			json.WriteByte(',')
		}
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
	for colNo := range columns {
		if colNo != 0 {
			s.Buffer.WriteRune(',')
		}
		s.Buffer.WriteString(fmt.Sprintf("\"%s\":", s.columnName(columns[colNo], colNo)))
		s.serializeValue(tuple.ColValue(tabdef.Order(colNo)), columns[colNo].Type)
	}
	s.Buffer.WriteRune('}')
}

func (s *resultJsonSerializer) serializeValue(value []byte, t hgtype.ColType) {
	switch t.Tag() {
	case rawtype.TypeStr:
		s.Buffer.WriteString(fmt.Sprintf("\"%s\"", string(rawtype.TypeOp.GetData(value, t.Args()))))
	case rawtype.TypeInt8:
		s.Buffer.WriteString(fmt.Sprintf("%d", bparse.Parse.Int8(value)))
	default:
		log.Panic("unsupported hgtype type: " + t.Tag().ToStr())
	}
}

func (s *resultJsonSerializer) columnName(column page.PatternCol, colNo int) string {
	if column.Name != "" {
		return column.Name
	} else {
		return fmt.Sprintf("col%d", colNo)
	}
}

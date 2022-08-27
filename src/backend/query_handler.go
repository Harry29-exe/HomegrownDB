package backend

import (
	"HomegrownDB/backend/executor/qrow"
	"HomegrownDB/backend/parser"
)

func HandleQuery(query string) (qrow.RowBuffer, error) {
	tree := parser.Parse(query)
}

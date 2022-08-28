package backend

import (
	"HomegrownDB/backend/executor/qrow"
	"HomegrownDB/backend/parser"
)

func HandleQuery(query string) (qrow.RowBuffer, error) {
	tree, err := parser.Parse(query)
	if err != nil {
		return nil, err
	}

}

package command

import "HomegrownDB/sql/parser"

func Handle(command string) {
	if command[0] == '.' {
		handleMeta(command)
	} else {
		sqlStatement := parser.ParseSql(command)
	}
}

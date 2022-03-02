package io

func HandleCommand(command string) {
	if command[0] == '.' {
		HandleMetaCommand(command)
	} else {
		//sqlStatement := parser.ParseSql(command)
	}
}

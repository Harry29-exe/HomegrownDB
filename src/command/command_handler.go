package command

import "HomegrownDB/sql"

func Handle(command string) {
	if (*command)[0] == '.' {
		handleMetaCommand(command)
	} else {
		delegateSqlCommand(command)
	}
}

func delegateSqlCommand(command *string) {
	sql.HandleSqlCommand(command)
}

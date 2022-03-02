package parser

type SqlCommandType uint

const (
	INSERT SqlCommandType = iota
	SELECT
	DELETE
	SET
)

type SqlCommand struct {
	commandType SqlCommandType
}

func ParseSql(sql string) SqlCommand {
	printType(DELETE)
	return SqlCommand{}
}

func printType(t SqlCommandType) {

}

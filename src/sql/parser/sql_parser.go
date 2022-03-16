package parser

type SqlKeyWord uint16

const (
	INSERT SqlKeyWord = iota
	SELECT
	DELETE
	UPDATE
	FROM
	WHERE
	JOIN
	SET
)

type SqlParseTree struct {
}

type SqlCommandParts struct {
}

func ParseSql(sql string) SqlParseTree {

	for _, char := range sql {
		switch {
		case char < 91 && char > 64:
			if
		}
	}
	printType(DELETE)
	return SqlParseTree{}
}

func startsWithKeyword(str string) {

}

func printType(t SqlCommandType) {

}

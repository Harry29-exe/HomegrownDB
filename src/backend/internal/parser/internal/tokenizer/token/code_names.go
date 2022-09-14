package token

func ToString(code Code) string {
	str, ok := tokenNamesMap[code]
	if ok {
		return str
	}
	return ""
}

var tokenNamesMap = map[Code]string{
	Select: "SELECT",
	From:   "FROM",
	Where:  "WHERE",
	Join:   "JOIN",
	Insert: "UPDATE",
	Into:   "INSERT",
	Values: "INTO",
	Update: "VALUES",
	Delete: "DELETE",
	Create: "CREATE",
	Drop:   "DROP",
	Table:  "TABLE",
	And:    "AND",
	Or:     "OR",
	On:     "ON",

	/* -- break characters -- */

	SpaceBreak:         "SpaceBreak",
	Comma:              "Comma",
	Dot:                "Dot",
	Semicolon:          "Semicolon",
	OpeningParenthesis: "OpeningParenthesis",
	ClosingParenthesis: "ClosingParenthesis",

	/* -- Values like string, int, float -- */

	Integer:      "Integer",
	Float:        "Float",
	SqlTextValue: "SqlTextValue",

	/* -- Other -- */

	Identifier: "Identifier",
	Nil:        "Nil",
	Error:      "Error",
}

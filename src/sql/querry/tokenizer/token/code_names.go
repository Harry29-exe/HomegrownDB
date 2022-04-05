package token

func ToString(code Code) string {
	str, ok := tokenNamesMap[code]
	if ok {
		return str
	}
	return ""
}

const (
	SelectStr = "SELECT"
	FromStr   = "FROM"
	WhereStr  = "WHERE"
	JoinStr   = "JOIN"
	UpdateStr = "UPDATE"
	DeleteStr = "DELETE"
	CreateStr = "CREATE"
	DropStr   = "DROP"
	TableStr  = "TABLE"
	AndStr    = "AND"
	OrStr     = "OR"
	OnStr     = "ON"

	/* -- break characters -- */

	SpaceBreakStr = "SpaceBreak"
	CommaStr      = "Comma"
	DotStr        = "Dot"
	SemicolonStr  = "Semicolon"

	/* -- Values like string, int, float -- */

	IntegerStr      = "Integer"
	FloatStr        = "Float"
	SqlTextValueStr = "SqlTextValue"

	/* -- Other -- */

	TextStr = "Text"
)

var tokenNamesMap = map[Code]string{
	Select: SelectStr,
	From:   FromStr,
	Where:  WhereStr,
	Join:   JoinStr,
	Update: UpdateStr,
	Delete: DeleteStr,
	Create: CreateStr,
	Drop:   DropStr,
	Table:  TableStr,
	And:    AndStr,
	Or:     OrStr,
	On:     OnStr,

	/* -- break characters -- */

	SpaceBreak: SpaceBreakStr,
	Comma:      CommaStr,
	Dot:        DotStr,
	Semicolon:  SemicolonStr,

	/* -- Values like string, int, float -- */

	Integer:      IntegerStr,
	Float:        FloatStr,
	SqlTextValue: SqlTextValueStr,

	/* -- Other -- */

	Text: TextStr,
}

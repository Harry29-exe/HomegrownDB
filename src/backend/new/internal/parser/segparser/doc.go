// Package segparser aggregate segment parsers.
// Each parser parses segment that it can understand often using
// other parsers. E.g. Select parser will use tables parser
// and tables parser will use table parser.
//
// Each parser should start at internal.TokenSource current token and
// return value when internal.TokenSource is pointing at last token parsed
// by parser. E.g. Table parser should start at table name and finish
// when internal.TokenSource point either to table name or table alias (if exist)
package segparser

import (
	"HomegrownDB/backend/new/internal/parser/tokenizer"
	"HomegrownDB/backend/new/internal/parser/validator"
)

type tkSource = tokenizer.TokenSource
type tkValidator = validator.Validator

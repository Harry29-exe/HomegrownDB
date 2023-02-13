// Package parse aggregate segment parsers.
// Each parser parses segment that it can understand often using
// other parsers. E.g. Select parser will use tables parser
// and tables parser will use tabdef parser.
//
// Each parser should start at internal.TokenSource current token and
// return value when internal.TokenSource is pointing at last token parsed
// by parser. E.g. Table parser should start at tabdef name and finish
// when internal.TokenSource point either to tabdef name or tabdef alias (if exist)
package parse

import (
	"HomegrownDB/backend/internal/parser/tokenizer"
	"HomegrownDB/backend/internal/parser/validator"
)

type tkSource = tokenizer.TokenSource
type tkValidator = validator.Validator

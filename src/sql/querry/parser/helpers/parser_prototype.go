package helpers

import "HomegrownDB/sql/querry/parser/def"

type ParserHelper struct {
	source def.TokenSource
}

func (h *ParserHelper) Init(source def.TokenSource) {
	h.source = source
}

package common

import (
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If either method reach array end it returns nil without changing pointer
type TokenSource interface {
	Next() tokenizer.Token      // Next move pointer forwards and returns
	Prev() tokenizer.Token      // Prev move pointer backwards and returns
	Current() tokenizer.Token   // Current returns token which has been returned with last method
	History() []tokenizer.Token // History returns all token from beginning to the one that Next would return

	Checkpoint() // Checkpoint creates new checkpoint for parser to rollback
	Commit()     // Commit deletes last checkpoint
	Rollback()   // Rollback to last checkpoint and removes this checkpoint
}

type Parser interface {
	Parse(source TokenSource) (parsetree.Node, error)
}

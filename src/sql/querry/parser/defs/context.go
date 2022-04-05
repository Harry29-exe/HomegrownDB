package defs

import (
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer/token"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If either method reach array end it returns nil without changing pointer
type TokenSource interface {
	Next() token.Token      // Next move pointer forwards and returns, if source has no more tokens it returns nil
	Prev() token.Token      // Prev move pointer backwards and returns, if source has no more tokens it returns nil
	Current() token.Token   // Current returns token which has been returned with last method
	History() []token.Token // History returns all token from beginning to the one that Next would return

	Checkpoint() // Checkpoint creates new checkpoint for parser to rollback
	Commit()     // Commit deletes last checkpoint
	Rollback()   // Rollback to last checkpoint and removes this checkpoint
}

type Parser interface {
	Parse(source TokenSource) (ptree.Node, error)
}

type ParserPrototype struct {
	Root ptree.Node
}

func (p *ParserPrototype) Attach(node ptree.Node, err error) error {
	if err != nil {
		return err
	}

	err = p.Root.AddChild(node)
	return err
}

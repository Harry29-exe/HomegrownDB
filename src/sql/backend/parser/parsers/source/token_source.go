package source

import (
	"HomegrownDB/sql/backend/tokenizer/token"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If method reach array's end it returns token with code token.Nil without changing pointer
type TokenSource interface {
	Next() token.Token      // Next move pointer forwards and returns, if source has no more tokens it returns nil
	Prev() token.Token      // Prev move pointer backwards and returns, if source has no more tokens it returns nil
	Current() token.Token   // Current returns token which has been returned with last method
	History() []token.Token // History returns all token from beginning to the one that Next would return

	Checkpoint() // Checkpoint creates new checkpoint for parser to rollback
	// CommitAndCheckpoint todo probably because of new aproach to helpers this hould be deleted
	CommitAndCheckpoint() // CommitAndCheckpoint commits and creates new checkpoint
	Commit()              // Commit deletes last checkpoint
	Rollback()            // Rollback to last checkpoint and removes this checkpoint
}

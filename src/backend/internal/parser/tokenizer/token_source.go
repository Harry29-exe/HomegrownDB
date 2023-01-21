package tokenizer

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If method reach array's end it returns token with code token.Nil without changing pointer
type TokenSource interface {
	Next() token.Token         // Next move pointer forwards and returns, if source has no more tokens it returns nil
	Prev() token.Token         // Prev move pointer backwards and returns, if source has no more tokens it returns nil
	Current() token.Token      // Current returns token which has been returned with last method
	CurrentTokenIndex() uint32 // CurrentTokenIndex returns index of token that source's pointer is pointing to
	History() []token.Token    // History returns all token from beginning to the one that Next would return

	Get(index uint) token.Token           // Get returns token with given index
	GetPtrRelative(index int) token.Token // GetPtrRelative returns token at index = ptr + index, so 0 means current, -1 means last, et cetera

	Checkpoint()                       // Checkpoint creates new checkpoint for segparser to rollback
	Commit()                           // Commit deletes last checkpoint
	CommitAndInitNode(node pnode.Node) // CommitAndInitNode invoke commit, and set start and end token of received test
	Rollback()                         // Rollback to last checkpoint and removes this checkpoint
}

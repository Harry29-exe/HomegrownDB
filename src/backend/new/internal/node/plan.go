package node

type Plan struct {
	targetList []TargetEntry
}

type TargetEntry struct {
	entryType *Expr // expression to evaluate to
	name      string
}

package analyser

import "HomegrownDB/backend/internal/analyser/internal"

type Tree struct {
	RootType rootType
	Ctx      *internal.AnalyserCtx
	Root     any
}

type rootType = uint8

const (
	RootTypeSelect rootType = iota
	RootTypeInsert
)

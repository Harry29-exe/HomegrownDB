package anode

import (
	"HomegrownDB/backend/internal/shared/qctx"
)

type Select struct {
	Tables []qctx.QTableId
	Fields SelectFields
}

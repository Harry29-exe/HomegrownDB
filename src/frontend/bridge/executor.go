package bridge

import (
	"HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

type Executor interface {
	Execute(query string, tx tx.Id, auth auth.Ctx) []page.RTuple
}

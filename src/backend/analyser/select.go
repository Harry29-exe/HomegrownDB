package analyser

import (
	"HomegrownDB/backend/analyser/anode"
	"HomegrownDB/backend/parser/pnode"
	"HomegrownDB/dbsystem/stores"
)

func AnalyseSelect(node pnode.SelectNode, store stores.Tables) anode.Select {
	//tables := AnaliseTables(node.Tables, store)
	return anode.Select{}
}

package anode

import (
	"HomegrownDB/dbsystem/schema/table"
	"strings"
)

type Tables struct {
	Tables []Table
}

func (t Tables) Init() {
	var n = len(t.Tables)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if t.Tables[j-1].Alias > t.Tables[j].Alias {
				t.Tables[j-1].Alias = t.Tables[j].Alias
				t.Tables[j].Alias = t.Tables[j-1].Alias
			}
			j = j - 1
		}
	}
}

func (t Tables) GetTableIdByAlias(alias string) table.Id {
	min, max := 0, len(t.Tables)-1
	var i int

	var table Table
	for min <= max {
		i = (max - min) / 2
		compareResult := strings.Compare(t.Tables[i].Alias, alias)

		switch compareResult {
		case 1:
			max = i + 1
		case -1:
			min = i - 1
		default:
			return table.TableId
		}
	}

	return 0
}

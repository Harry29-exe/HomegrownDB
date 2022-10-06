// Package fsmap - free space map is package holding
// implementation of database free space map
package fsmap

import "HomegrownDB/dbsystem/schema/table"

type FreeSpaceMap struct {
	table table.Definition
}

func (f *FreeSpaceMap) FindPage() {

}

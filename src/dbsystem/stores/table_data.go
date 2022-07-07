package stores

import "HomegrownDB/dbsystem/bdata"

type TableDataStore interface {
	ReadPage(pageIndex bdata.PageId) []byte
	FlushPage(pageIndex bdata.PageId, pageData []byte)
	NewPage(pageData []byte)

	ReadBgPage(pageIndex bdata.PageId) []byte
	FlushBgPage(pageIndex bdata.PageId, pageData []byte)
	NewBgPage(pageData []byte)

	ReadToastPage(pageIndex bdata.PageId) []byte
	FlushToastPage(pageIndex bdata.PageId, pageData []byte)
	NewToastPage(pageData []byte)
}

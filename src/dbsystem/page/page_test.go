package page_test

import "HomegrownDB/dbsystem/page"

var pUtils pageUtils = pageUtils{}

type pageUtils struct{}

func (u pageUtils) EmptyPage() page.Page {

}

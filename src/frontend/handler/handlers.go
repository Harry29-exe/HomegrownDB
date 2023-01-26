package handler

import "HomegrownDB/dbsystem/hg/di"

type Handlers struct {
	SqlHandler SqlHandler
}

func NewHandlers(container di.FrontendContainer) Handlers {

}

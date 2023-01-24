package server

import (
	"HomegrownDB/frontend/handler"
	"HomegrownDB/frontend/server/http"
)

type DBServer interface {
	Start() error
	Stop() error
}

func CreateDefaultServer(address string, port string, handlers handler.Handlers) DBServer {
	return http.CreateDBServer(address, port, handlers)
}

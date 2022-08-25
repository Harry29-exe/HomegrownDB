package server

import (
	"HomegrownDB/frontend/server/http"
)

type DBServer interface {
	Start() error
	Stop() error
}

func CreateDefaultService(address string, port string) DBServer {
	return http.CreateDBServer(address, port)
}

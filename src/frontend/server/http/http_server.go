package http

import "github.com/gofiber/fiber/v2"

type DBServer struct {
	app             *fiber.App
	addressWithPort string
}

func (s *DBServer) Start() error {
	return s.app.Listen(s.addressWithPort)
}

func (s *DBServer) Stop() error {
	return s.app.Shutdown()
}

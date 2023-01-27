package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DBServer struct {
	app             *fiber.App
	addressWithPort string
}

func (s *DBServer) Start() error {
	errChan := make(chan error)
	testChan := make(chan error)
	go func() {
		errChan <- s.app.Listen(s.addressWithPort)
	}()
	go testServer(s.app, testChan)

	select {
	case err := <-errChan:
		return err
	case result := <-testChan:
		return result
	}
}

func (s *DBServer) Stop() error {
	return s.app.Shutdown()
}

func testServer(app *fiber.App, testChan chan error) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		testChan <- err
		return
	}
	_, err = app.Test(request, 10_000)
	if err != nil {
		testChan <- err
		return
	}
	testChan <- nil
	return
}

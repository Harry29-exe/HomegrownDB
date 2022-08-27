package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
)

func CreateDBServer(address, port string) *DBServer {
	app := fiber.New()

	app.Post("/api/", func(ctx *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello")
		return ctx.SendString(msg)
	})

	app.Post("/api/:hello_msg", func(ctx *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello %s", ctx.Params("hello_msg"))
		return ctx.SendString(msg)
	})
	app.Post("/api/large", func(ctx *fiber.Ctx) error {
		return ctx.SendStream(io.Reader())
	})

	return &DBServer{
		app:             app,
		addressWithPort: address + ":" + port,
	}
}

func registerSendQueryEndpoint(app *fiber.App) {
	app.Post("/db/query")
}

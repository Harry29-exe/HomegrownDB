package http

import (
	auth "HomegrownDB/dbsystem/auth"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/frontend/handler"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateDBServer(address, port string, handlers handler.Handlers) *DBServer {
	app := fiber.New()

	app.Post("/api/test", func(ctx *fiber.Ctx) error {
		msg := fmt.Sprintf("Test successful")
		return ctx.SendString(msg)
	})

	registerSendQueryEndpoint(app, handlers)

	return &DBServer{
		app:             app,
		addressWithPort: address + ":" + port,
	}
}

func registerSendQueryEndpoint(app *fiber.App, handlers handler.Handlers) {
	app.Post("/db-api/query", func(ctx *fiber.Ctx) error {
		body := &SqlRequestBody{}
		err := json.Unmarshal(ctx.Body(), body)
		if err != nil {
			return err
		}
		query := body.Query
		txId := tx.InvalidId
		if body.TxId != nil {
			txId = *body.TxId
		}
		var authData auth.Authentication = nil

		result, err := handlers.SqlHandler.Handle(query, txId, authData)
		if err != nil {
			return err
		}
		return ctx.Send(result.Result())
	})
}

type SqlRequestBody struct {
	TxId  *tx.Id
	Query string
}

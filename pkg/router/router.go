package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucaslimafernandes/rinha-backend-2025/pkg/router/handler"
)

// Endpoint	Descrição
// POST /payments	Intermedia a requisição para o processamento dum pagamento.
// GET /payments-summary	Exibe detalhes das requisições de processamento de pagamentos.

func SetupRoutes(app *fiber.App) {

	// Rotas
	api := app.Group("/")
	api.Post("payments", handler.ApiPayments)
	api.Get("payments-summary", handler.ApiPayments)

}

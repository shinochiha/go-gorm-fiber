package handler

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/model"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func CreateOrderHandler(ctx *fiber.Ctx) error {
	// 1. Initiate coreapi client
	c := coreapi.Client{}
	c.New("SB-Mid-server-KTQQ6LNGHzxPixnCuMzqr118", midtrans.Sandbox)
	// Parse request body
	var request model.OrderRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	/// Create Midtrans transaction request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.OrderID,
			GrossAmt: int64(request.Amount),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "ITEM01",
				Price: int64(request.Amount),
				Qty:   1,
				Name:  request.Description,
			},
		},
	}

	// Create Midtrans transaction using Core API
	resp, err := c.ChargeTransaction(chargeReq)
	if err != nil {
		log.Println("Midtrans transaction error:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

	// Save the transaction to the database
	log.Println(resp)
	log.Println("--------------------")
	transaction := model.Transaction{
		OrderID:            request.OrderID,
		TransactionDetails: resp.RedirectURL,
	}
	if err := config.DB.Create(&transaction).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transaction",
		})
	}

	response := model.OrderResponse{
		RedirectURL: resp.RedirectURL,
	}

	return ctx.JSON(response)

}

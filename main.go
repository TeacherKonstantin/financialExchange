package main

import (
	"financialExchange/storage"
	"financialExchange/enum"
	"financialExchange/modifiers"
	"fmt"
	"log/slog"
	"os"
	"uuid"
	"time"
)

func main(){
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	

	orderBook := storage.NewStore[uuid.UUID, enum.Order](10)

	uuidOrder1 := uuid.NewV7()
	uuidOrder2 := uuid.NewV7()

	now := time.Now().UTC()

	orderBuy := 	enum.Order{
						ID: uuidOrder1,
						TraderID: uuid.NewV4(),
						Side: enum.SideBuy,
						Type: enum.OrderTypeLimit,
						Symbo: "USD",
						Price: 1000.00,
						Quantity: 2.00,
						Filled: 0.00,
						Status: enum.OrderStatusNew,
						CreatedAt: now,
						UpdatedAt: now,
							}

	

	orderSell := 	enum.Order{
						ID: uuidOrder2,
						TraderID: uuid.NewV4(),
						Side: enum.SideSell,
						Type: enum.OrderTypeLimit,
						Symbo: "USD",
						Price: 1000.00,
						Quantity: 1.00,
						Filled: 0.00,
						Status: enum.OrderStatusNew,
						CreatedAt: now,
						UpdatedAt: now,
						}

	fmt.Scanln()

	logger.Info("Заявки созданы", "buy_order", orderBuy, "sell_order", orderSell)
	orderBook.Put(uuidOrder1, orderBuy)
	orderBook.Put(uuidOrder2, orderSell)

	fmt.Scanln()

	orderBook.Update(uuidOrder1,  modifiers.ModifierForOrder(orderSell.Quantity))
	orderBook.Update(uuidOrder2,  modifiers.ModifierForOrder(orderBuy.Quantity))

	
	fmt.Scanln()

}
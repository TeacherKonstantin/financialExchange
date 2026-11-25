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

	matchOrders(orderBook, orderBuy.ID, orderSell.ID, logger)

	// orderBook.Update(uuidOrder1,  modifiers.ModifierForOrder(orderSell.Quantity))
	// orderBook.Update(uuidOrder2,  modifiers.ModifierForOrder(orderBuy.Quantity))

	
	fmt.Scanln()

}

func matchOrders(storage *storage.Store[uuid.UUID, enum.Order], buyID, sellID uuid.UUID, logger *slog.Logger){
	buyOrder, buyOk := storage.Get(buyID)
	sellOrder, sellOk := storage.Get(sellID)

	if !buyOk || !sellOk{
		logger.Error("Возникли ошибки с поиском заказов в хранилище")
		return
	}

	if buyOrder.Symbo != sellOrder.Symbo{
		logger.Warn("Сделка не может быть осуществлена в связи несовпадением валют")
		return
	}
	
	if buyOrder.Status == enum.OrderStatusFilled || sellOrder.Status == enum.OrderStatusFilled{
		logger.Warn("Заказы уже завершены")
		return
	}

	if buyOrder.Price < sellOrder.Price{
		logger.Info("Сделка не может быть совершена из-за несовпадения цен")
		return
	}

	quantityTrade := min(buyOrder.Remaining(), sellOrder.Remaining())

	updatedBuyOrder, ok:= storage.Update(buyID, modifiers.ModifierForOrder(quantityTrade))

	if !ok{
		logger.Error("Произошла ошибка при обновлении buyOrder")
	}

	updatedSellOrder, ok:= storage.Update(sellID, modifiers.ModifierForOrder(quantityTrade))

	if !ok{
		logger.Error("Произошла ошибка при обновлении sellOrder")
	}

	logger.Info("Заявки обнавлены", "buy_order", updatedBuyOrder, "sell_order", updatedSellOrder)

	logger.Info("Сделка успешно проведена",
				slog.Float64("quantity", quantityTrade),
				slog.Any("buy_order", updatedBuyOrder),
				slog.Any("sell_order", updatedSellOrder),
	)	

}
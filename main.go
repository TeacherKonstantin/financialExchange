package main

import (
	"financialExchange/storage"
	"fmt"
	"log/slog"
	"os"
	"uuid"
)

func main(){
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	// currencies := storage.NewStorage[string](5)
	// currencies.Add("BTC-RUB")
	// currencies.Add("ETH-RUB")
	// currencies.Add("USD-RUB")
	// currencies.Add("USD-RUB")

	// lengthOfCurrencies := currencies.Length()
	// logger.Info("Работа с валютами:",
	//  slog.Int("Количество уникальных комбинаций", lengthOfCurrencies))

	traderStore := storage.NewStore[uuid.UUID, float64](10)
	traderID := uuid.NewV7()

	traderStore.Put(traderID, 50000.0004)
	logger.Info("Новый трейдер добавлен в хранилище:",
	 slog.String("UUID", traderID.String()),
	slog.Float64("Начальный баланс", 50000.0004))


	if balance, ok := traderStore.Get(traderID); ok{
		logger.Info("поиск в хранилище:",
	 	slog.String("найден трейдер с UUID", traderID.String()),
		slog.Float64("его текущий баланс", balance))
	}else{
		logger.Error("элемент не найден в хранилище:",
	 	slog.String("проверьте UUID", traderID.String()),)
	}


	fmt.Scanln()

}
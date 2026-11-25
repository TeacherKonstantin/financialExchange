package enum

import (
	//"google/uuid"
	"log/slog"
	"math"
	"time"
	"uuid"
)

type Side int8

const (
	SideBuy Side = iota + 1
	SideSell
)

//Stringer type

func (side Side) String() string {
	switch side {
	case SideBuy: return "Покупка"
	case SideSell: return "Продажа"
	default: return "Операция не определена"
	}
}

type OrderType int8
const (
	OrderTypeLimit OrderType = iota + 1 //Заявка по фиксированной цене
	OrderTypeMarket						//Заявка по рыночной цене
)

func (t OrderType) String() string {
	switch t {
	case OrderTypeLimit: return "Заявка на фиксированную цену"
	case OrderTypeMarket: return "Заявка по рыночной стоимости"
	default: return "Тип заявки не определен"
	}
}

type OrderStatus int8

const (
	OrderStatusNew OrderStatus = iota + 1	// Зарегистрирована на бирже
	OrderStatusPartiallyFilled				// Частично выполнена
	OrderStatusFilled						// Выполнена
	OrderStatusCancelled					// Отмена трейдером
	OrderStatusRejected						// Отклонена из за нарушения правил или нехватки суммы для перевода
)

func (status OrderStatus) String() string {
	switch status {
	case OrderStatusNew: return "Заявка зарегистрирована"
	case OrderStatusPartiallyFilled: return "Заявка частично выполнена"
	case OrderStatusFilled: return "Заявка выполнена"
	case OrderStatusCancelled: return "Трейдер отменил заявку"
	case OrderStatusRejected: return "Заявка заморожена биржей"
	default: return "Статус не определяется"
	}
}

type Trader struct {
	ID uuid.UUID
	Name string
	Cash float64
	CreatedAt time.Time
}

type Order struct {
	ID uuid.UUID
	TraderID uuid.UUID
	Side Side
	Type OrderType
	Symbo string
	Price float64
	Quantity float64
	Filled float64
	Status OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

//интерфейс slog.LogValuer
func (order Order)Logvalue() slog.Value{
	return slog.GroupValue(
			slog.String("ID: ", order.ID.String()),
			slog.String("Side: ", order.Side.String()),
			slog.String("Status: ", order.Status.String()),
			slog.Float64("Value:", order.Price),
			slog.Float64("Quantity: ", order.Quantity),
			slog.Float64("Filled: ", order.Filled),
			slog.String("Symbo: ", order.Symbo),
			slog.String("TraderID: ", order.TraderID.String()),
			slog.String("Updated at: ", order.UpdatedAt.String()))
}

func(order *Order) Fill(count float64){
		order.Filled += count
		order.UpdatedAt = time.Now().UTC()
//в будущем учесть погрешности 
// которые возникают с float64 при арифметических операциях
		if order.Filled >= order.Quantity{
			order.Status = OrderStatusFilled
		} else{
			order.Status = OrderStatusPartiallyFilled
		}
}

func (order *Order) Remaining() float64 {
	return order.Quantity - order.Filled
}

func (order *Order) IsActive() bool {
	return order.Status == OrderStatusNew || order.Status == OrderStatusPartiallyFilled
}

type Trade struct { 
	ID 			uuid.UUID
	Symbol 		string
	BuyOrderID 	uuid.UUID
	SellOrderID uuid.UUID
	BuyerID 	uuid.UUID
	SellerID 	uuid.UUID
	Price 		float64
	Quantity 	float64
	ExecutedAt 	time.Time
}

type Position struct {
	TraderID 	uuid.UUID
	Symbol   	string
	Quantity 	float64
	AvgPrice 	float64
	Realized 	float64
}

func RoundPrice(price float64) float64 { 
	return math.Round(price*100) / 100
}

package dto

import (
	"time"

	"WB/internal/order/model"
)

type OrderDB struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string // Update this to ShardKey
	SmID              int
	DateCreated       time.Time
	OofShard          string
	Delivery          DeliveryDB
	Payment           PaymentDB
	Items             []ItemDB
}

type DeliveryDB struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type PaymentDB struct {
	Transaction  string
	RequestID    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDt    int64
	Bank         string
	DeliveryCost int
	GoodsTotal   int
	CustomFee    int
}

type ItemDB struct {
	ChrtID      int
	TrackNumber string
	Price       int
	Rid         string
	Name        string
	Sale        int
	Size        string
	TotalPrice  int
	NmID        int
	Brand       string
	Status      int
}

func NewOrderDB(order model.Order) OrderDB {
	return OrderDB{
		OrderUID:          order.OrderUID,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.Shardkey,
		SmID:              order.SmID,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
		Delivery: DeliveryDB{
			Name:    order.Delivery.Name,
			Phone:   order.Delivery.Phone,
			Zip:     order.Delivery.Zip,
			City:    order.Delivery.City,
			Address: order.Delivery.Address,
			Region:  order.Delivery.Region,
			Email:   order.Delivery.Email,
		},
		Payment: PaymentDB{
			Transaction:  order.Payment.Transaction,
			RequestID:    order.Payment.RequestID,
			Currency:     order.Payment.Currency,
			Provider:     order.Payment.Provider,
			Amount:       order.Payment.Amount,
			PaymentDt:    int64(order.Payment.PaymentDt),
			Bank:         order.Payment.Bank,
			DeliveryCost: order.Payment.DeliveryCost,
			GoodsTotal:   order.Payment.GoodsTotal,
			CustomFee:    order.Payment.CustomFee,
		},
		Items: func(items []model.Item) []ItemDB {
			itemsDB := make([]ItemDB, len(items))
			for i, item := range items {
				itemsDB[i] = ItemDB{
					ChrtID:      item.ChrtID,
					TrackNumber: item.TrackNumber,
					Price:       item.Price,
					Rid:         item.Rid,
					Name:        item.Name,
					Sale:        item.Sale,
					Size:        item.Size,
					TotalPrice:  item.TotalPrice,
					NmID:        item.NmID,
					Brand:       item.Brand,
					Status:      item.Status,
				}
			}
			return itemsDB
		}(order.Items),
	}
}

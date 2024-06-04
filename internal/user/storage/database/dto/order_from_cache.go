package dto

import (
	"WB/internal/order/model"
	"time"
)

type OrderFromCache struct {
	OrderID         string   `json:"order_id"`
	TrackNumber     string   `json:"track_number"`
	Entry           string   `json:"entry"`
	Locale          string   `json:"locale"`
	Delivery        Delivery `json:"delivery"`
	Payment         Payment  `json:"payment"`
	Items           []Item   `json:"items"`
	CustomerID      string   `json:"customer_id"`
	DeliveryService string   `json:"delivery_service"`
	ShardKey        string   `json:"shardkey"`
	SmID            int      `json:"sm_id"`
	DateCreated     string   `json:"date_created"`
	OofShard        string   `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int64  `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func ConvertToOrderFromCache(order model.Order) *OrderFromCache {
	return &OrderFromCache{
		OrderID:         order.OrderUID,
		TrackNumber:     order.TrackNumber,
		Entry:           order.Entry,
		Locale:          order.Locale,
		Delivery:        ConvertDelivery(order.Delivery),
		Payment:         ConvertPayment(order.Payment),
		Items:           ConvertItems(order.Items),
		CustomerID:      order.CustomerID,
		DeliveryService: order.DeliveryService,
		ShardKey:        order.Shardkey,
		SmID:            order.SmID,
		DateCreated:     order.DateCreated.Format(time.RFC3339),
		OofShard:        order.OofShard,
	}
}

func ConvertDelivery(delivery model.Delivery) Delivery {
	return Delivery{
		Name:    delivery.Name,
		Phone:   delivery.Phone,
		Zip:     delivery.Zip,
		City:    delivery.City,
		Address: delivery.Address,
		Region:  delivery.Region,
		Email:   delivery.Email,
	}
}

func ConvertPayment(payment model.Payment) Payment {
	return Payment{
		Transaction:  payment.Transaction,
		RequestID:    payment.RequestID,
		Currency:     payment.Currency,
		Provider:     payment.Provider,
		Amount:       payment.Amount,
		PaymentDt:    int64(payment.PaymentDt),
		Bank:         payment.Bank,
		DeliveryCost: payment.DeliveryCost,
		GoodsTotal:   payment.GoodsTotal,
		CustomFee:    payment.CustomFee,
	}
}

func ConvertItems(items []model.Item) []Item {
	result := make([]Item, len(items))
	for i, item := range items {
		result[i] = Item{
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
	return result
}

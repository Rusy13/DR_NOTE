package dto

import (
	"WB/internal/order/model"
	"time"
)

func ConvertToOrderCache(order model.Order) *OrderFromCache {
	return &OrderFromCache{
		OrderID:         order.OrderUID,
		TrackNumber:     order.TrackNumber,
		Entry:           order.Entry,
		Locale:          order.Locale,
		Delivery:        ConvertDeliveryC(order.Delivery),
		Payment:         ConvertPaymentC(order.Payment),
		Items:           ConvertItemsC(order.Items),
		CustomerID:      order.CustomerID,
		DeliveryService: order.DeliveryService,
		ShardKey:        order.Shardkey,
		SmID:            order.SmID,
		DateCreated:     order.DateCreated.Format(time.RFC3339),
		OofShard:        order.OofShard,
	}
}

func ConvertDeliveryC(delivery model.Delivery) Delivery {
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

func ConvertPaymentC(payment model.Payment) Payment {
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

func ConvertItemsC(items []model.Item) []Item {
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

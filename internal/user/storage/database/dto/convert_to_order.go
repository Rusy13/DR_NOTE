package dto

import "WB/internal/order/model"

func ConvertToOrder(orderDB OrderDB) model.Order {
	items := make([]model.Item, len(orderDB.Items))
	for i, itemDB := range orderDB.Items {
		items[i] = model.Item{
			ChrtID:      itemDB.ChrtID,
			TrackNumber: itemDB.TrackNumber,
			Price:       itemDB.Price,
			Rid:         itemDB.Rid,
			Name:        itemDB.Name,
			Sale:        itemDB.Sale,
			Size:        itemDB.Size,
			TotalPrice:  itemDB.TotalPrice,
			NmID:        itemDB.NmID,
			Brand:       itemDB.Brand,
			Status:      itemDB.Status,
		}
	}

	return model.Order{
		OrderUID:          orderDB.OrderUID,
		TrackNumber:       orderDB.TrackNumber,
		Entry:             orderDB.Entry,
		Locale:            orderDB.Locale,
		InternalSignature: orderDB.InternalSignature,
		CustomerID:        orderDB.CustomerID,
		DeliveryService:   orderDB.DeliveryService,
		Shardkey:          orderDB.ShardKey,
		SmID:              orderDB.SmID,
		DateCreated:       orderDB.DateCreated,
		OofShard:          orderDB.OofShard,
		Delivery: model.Delivery{
			Name:    orderDB.Delivery.Name,
			Phone:   orderDB.Delivery.Phone,
			Zip:     orderDB.Delivery.Zip,
			City:    orderDB.Delivery.City,
			Address: orderDB.Delivery.Address,
			Region:  orderDB.Delivery.Region,
			Email:   orderDB.Delivery.Email,
		},
		Payment: model.Payment{
			Transaction:  orderDB.Payment.Transaction,
			RequestID:    orderDB.Payment.RequestID,
			Currency:     orderDB.Payment.Currency,
			Provider:     orderDB.Payment.Provider,
			Amount:       orderDB.Payment.Amount,
			PaymentDt:    int(orderDB.Payment.PaymentDt),
			Bank:         orderDB.Payment.Bank,
			DeliveryCost: orderDB.Payment.DeliveryCost,
			GoodsTotal:   orderDB.Payment.GoodsTotal,
			CustomFee:    orderDB.Payment.CustomFee,
		},
		Items: items,
	}
}

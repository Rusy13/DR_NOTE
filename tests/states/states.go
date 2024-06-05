package states

import (
	"WB/internal/order/model"
	"time"
)

const (
	OrderUID1          string = "b563feb7b2b84b6test4"
	Track_number       string = "WBILMTESTTRACK"
	Entry              string = "WBIL"
	Locale             string = "en"
	Internal_signature string = ""
	Customer_id        string = "test"
	Delivery_service   string = "meest"
	Shardkey           string = "9"
	Sm_id              int    = 99
	Oof_shard          string = "1"
)

var (
	Date_created = time.Now()

	Delivery1 = model.Delivery{
		Name:    "John Doe",
		Phone:   "123456789",
		Zip:     "12345",
		City:    "New York",
		Address: "123 Main St",
		Region:  "NY",
		Email:   "john@example.com",
	}

	Payment1 = model.Payment{
		Transaction:  "123456",
		RequestID:    "789012",
		Currency:     "USD",
		Provider:     "PayPal",
		Amount:       1000,
		PaymentDt:    1637907727,
		Bank:         "Bank of America",
		DeliveryCost: 50,
		GoodsTotal:   950,
		CustomFee:    0,
	}

	Item1 = model.Item{
		ChrtID:     123456,
		Name:       "Item 1",
		Price:      100,
		Size:       "Medium",
		TotalPrice: 100,
		Status:     1,
	}

	SLItems = []model.Item{
		Item1,
	}
)

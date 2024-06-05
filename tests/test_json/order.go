package test_json

import (
	"fmt"
	"time"
)

// Order содержит JSON-объект заказа
const Order = `
{
  "order_uid": "b563feb7b2b84b6test100",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery_name": "Test Testov",
  "delivery_phone": "+9720000000",
  "delivery_zip": "2639809",
  "delivery_city": "Kiryat Mozkin",
  "delivery_address": "Ploshad Mira 15",
  "delivery_region": "Kraiot",
  "delivery_email": "test@gmail.com",
  "payment_transaction": "b563feb7b2b84b6test",
  "request_id": "",
  "currency": "USD",
  "provider": "wbpay",
  "amount": 1817,
  "payment_dt": 1637907727,
  "bank": "alpha",
  "delivery_cost": 1500,
  "goods_total": 317,
  "custom_fee": 1,
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`

// ExpectedOrder1 содержит ожидаемый JSON-объект заказа
var ExpectedOrder1 = `
{
    "order_id": "b563feb7b2b84b6test100",
    "track_number": "WBILMTESTTRACK",
    "entry": "WBIL",
    "locale": "en",
    "delivery": {
        "name": "Test Testov",
        "phone": "+9720000000",
        "zip": "2639809",
        "city": "Kiryat Mozkin",
        "address": "Ploshad Mira 15",
        "region": "Kraiot",
        "email": "test@gmail.com"
    },
    "payment": {
        "transaction": "b563feb7b2b84b6test",
        "request_id": "",
        "currency": "USD",
        "provider": "wbpay",
        "amount": 1817,
        "payment_dt": 1637907727,
        "bank": "alpha",
        "delivery_cost": 1500,
        "goods_total": 317,
        "custom_fee": 1
    },
    "items": [
        {
            "chrt_id": 9934930,
            "track_number": "WBILMTESTTRACK",
            "price": 453,
            "rid": "ab4219087a764ae0btest",
            "name": "Mascaras",
            "sale": 30,
            "size": "0",
            "total_price": 317,
            "nm_id": 2389212,
            "brand": "Vivienne Sabo",
            "status": 202
        }
    ],
    "customer_id": "test",
    "delivery_service": "meest",
    "shardkey": "9",
    "sm_id": 99,
    "date_created": "%s",
    "oof_shard": "1"
}
`

func init() {
	currentTime := time.Now().Add(0 * time.Hour)

	formattedTime := currentTime.Format("2006-01-02T03:04:05-07:00")

	ExpectedOrder1 = fmt.Sprintf(ExpectedOrder1, formattedTime)
}

// OrderAddNew содержит JSON-объект нового заказа
const OrderAddNew = `
{
  "order_uid": "b563feb7b2b84b6test100",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery_name": "Test Testov",
  "delivery_phone": "+9720000000",
  "delivery_zip": "2639809",
  "delivery_city": "Kiryat Mozkin",
  "delivery_address": "Ploshad Mira 15",
  "delivery_region": "Kraiot",
  "delivery_email": "test@gmail.com",
  "payment_transaction": "b563feb7b2b84b6test",
  "request_id": "",
  "currency": "USD",
  "provider": "wbpay",
  "amount": 1817,
  "payment_dt": 1637907727,
  "bank": "alpha",
  "delivery_cost": 1500,
  "goods_total": 317,
  "custom_fee": 1,
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`

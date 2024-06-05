//go:build integration
// +build integration

package order

import (
	"context"
	"fmt"
	"testing"

	"WB/internal/infrastructure/database/postgres/database"
	"WB/internal/order/model"
	"WB/tests/fixtures"
	"github.com/stretchr/testify/require"
)

func setUp(t *testing.T, db database.Database, tableNames []string) {
	t.Helper()
	ctx := context.Background()

	for _, tn := range tableNames {
		err := deleteData(ctx, db, tn)
		require.NoError(t, err)
	}
}

func deleteData(ctx context.Context, db database.Database, tableName string) error {
	_, err := db.Exec(ctx, fmt.Sprintf("DELETE FROM %s", tableName))
	return err
}

func fillDataBase(t *testing.T, db database.Database) {
	t.Helper()

	fillOrders(t, db)
	fillPayment(t, db)
	fillItems(t, db)
	fillDelivery(t, db)
}

func fillOrders(t *testing.T, db database.Database) {
	t.Helper()
	insertOrder(t, db, fixtures.Order().Valid1().Val())
}

func fillPayment(t *testing.T, db database.Database) {
	t.Helper()
	insertPayment(t, db, fixtures.Payment().Valid1().Val())
}

func fillItems(t *testing.T, db database.Database) {
	t.Helper()
	insertItem(t, db, fixtures.Item().Valid1().Val())
}

func fillDelivery(t *testing.T, db database.Database) {
	t.Helper()
	insertDelivery(t, db, fixtures.Delivery().Valid1().Val())
}

func insertOrder(t *testing.T, db database.Database, order model.Order) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	require.NoError(t, err)
}

func insertPayment(t *testing.T, db database.Database, payment model.Payment) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		payment.OrderUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	require.NoError(t, err)
}

func insertItem(t *testing.T, db database.Database, item model.Item) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		item.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
	require.NoError(t, err)
}

func insertDelivery(t *testing.T, db database.Database, delivery model.Delivery) {
	t.Helper()
	ctx := context.Background()
	_, err := db.Exec(ctx,
		`INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		delivery.OrderUID, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	require.NoError(t, err)
}

//go:build integration
// +build integration

package order

import (
	"WB/internal/infrastructure/database/postgres/database"
	redisApp "WB/internal/infrastructure/database/redis"
	"WB/internal/middleware"
	"WB/internal/order/delivery"
	serviceOrder "WB/internal/order/service"
	storageOrder "WB/internal/order/storage/database"
	"WB/internal/routes"
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

type OrderTestFixtures struct {
	del       *delivery.OrderDelivery
	mw        *middleware.Middleware
	db        *database.PGDatabase
	redisConn redis.Conn
	cancel    context.CancelFunc
	app       *mux.Router
}

func newOrderTestFixtures(t *testing.T) OrderTestFixtures {
	t.Helper()

	logger := zap.NewNop().Sugar()

	ctx, cancel := context.WithCancel(context.Background())
	db, err := database.New(ctx)
	require.NoError(t, err)

	redisConn, err := redisApp.Init()
	require.NoError(t, err)

	stOrder := storageOrder.New(db, redisConn, logger)
	svOrder := serviceOrder.New(stOrder)
	d := delivery.New(svOrder, logger)

	mw := middleware.New(logger)

	router := routes.GetRouter(d, mw)

	return OrderTestFixtures{
		del:       d,
		mw:        mw,
		db:        db,
		redisConn: redisConn,
		cancel:    cancel,
		app:       router,
	}
}

func (b *OrderTestFixtures) Close(t *testing.T) {
	b.cancel()
	err := b.db.Close()
	require.NoError(t, err)
	require.NoError(t, err)
	err = b.redisConn.Close()

}

//go:build integration
// +build integration

package order

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"WB/tests/test_json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	tableName1 = "items"
	tableName2 = "delivery"
	tableName3 = "payment"
	tableName4 = "orders"
)

//tableName1 = "orders"
//tableName2 = "payment"
//tableName3 = "items"
//tableName4 = "delivery"

var tableNames = []string{tableName1, tableName2, tableName3, tableName4}

func TestAddOrder(t *testing.T) {
	tf := newOrderTestFixtures(t)
	defer tf.Close(t)

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPost, "/order", strings.NewReader(test_json.OrderAddNew))
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.app.ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.JSONEq(t, `{"order_id":"b563feb7b2b84b6test100"}`, string(body))
	})
}

//go:build integration
// +build integration

package order

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"WB/tests/test_json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserOrder(t *testing.T) {
	tf := newOrderTestFixtures(t)
	defer tf.Close(t)

	t.Run("order not found", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/order/b563feb7b2b84b6test19", nil)
		request.Header.Set("Token", "user_token")
		respWriter := httptest.NewRecorder()

		tf.app.ServeHTTP(respWriter, request)
		resp := respWriter.Result()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodGet, "/order/b563feb7b2b84b6test100", nil)
		request.Header.Set("Token", "user_token")
		respWriter := httptest.NewRecorder()

		tf.app.ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, test_json.ExpectedOrder1, string(body))
	})
}

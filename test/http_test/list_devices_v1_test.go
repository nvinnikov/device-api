//go:build httptest || unit
// +build httptest unit

package http_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/assert"
	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
)

func TestListDevices(t *testing.T) {
	type testCasePositive struct {
		Page    string
		PerPage string
	}
	testsListDevicesPositive := []testCasePositive{
		{"1", "10"},
		{"1", "20"},
		{"2", "10"},
	}
	testsListDevicesNegative := []testCasePositive{
		{"0", "10"},
	}
	for _, tc := range testsListDevicesPositive {

		t.Run("Get devices", func(t *testing.T) {

			// Arrange
			var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			opts := url.Values{}
			opts.Add("page", tc.Page)
			opts.Add("perPage", tc.PerPage)
			ctx := context.Background()

			// Act
			items, _, err := client.ListDevices(ctx, opts)

			// Assert
			assert.NoError(t, err)
			assert.GreaterOrEqual(t, len(items.Items), 1)
		})
	}
	for _, tc := range testsListDevicesNegative {

		t.Run("Get devices", func(t *testing.T) {

			// Arrange
			var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			opts := url.Values{}
			opts.Add("page", tc.Page)
			opts.Add("perPage", tc.PerPage)
			ctx := context.Background()

			// Act
			_, _, err := client.ListDevices(ctx, opts)

			// Assert
			require.NotNil(t, err)
		})
	}
}

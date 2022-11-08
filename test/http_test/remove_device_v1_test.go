//go:build httptest1
// +build httptest1

package http_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/assert"
	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/models"
)

func TestDeleteDevice(t *testing.T) {

	type testCasePositive struct {
		Platform string
		UserId   string
	}
	testsCreateDevicePositive := []testCasePositive{
		{"Ios", "5342"},
		{"Android", "35456435"},
		{"Ubuntu", "4353452"},
	}
	var URL = envy.Get("BASE_URL", "http://localhost:8080")
	for _, tc := range testsCreateDevicePositive {
		t.Run("Delete device", func(t *testing.T) {
			// Arrange
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			device := models.CreateDeviceRequest{
				Platform: tc.Platform,
				UserID:   tc.UserId,
			}
			ctx := context.Background()

			// Act
			id, _, err := client.CreateDevice(ctx, device)
			assert.NoError(t, err)
			deletedDevice, _, err := client.RemoveDevice(ctx, strconv.Itoa(id.DeviceID))
			assert.NoError(t, err)

			// Assert
			assert.True(t, deletedDevice.Found)
		})
	}
}

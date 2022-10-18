package http_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/models"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/assert"
)

func TestDescribeDevice(t *testing.T) {

	type testCasePositive struct {
		Platform string
		UserId   string
	}
	testsDescribeDevicePositive := []testCasePositive{
		{"Ios", "111"},
		{"Android", "999"},
		{"Ubuntu", "555"},
	}
	var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")
	for _, tc := range testsDescribeDevicePositive {

		t.Run("Describe device", func(t *testing.T) {
			// Arrange
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			deviceCreate := models.CreateDeviceRequest{
				Platform: tc.Platform,
				UserID:   tc.UserId,
			}
			ctx := context.Background()

			// Act
			id, _, _ := client.CreateDevice(ctx, deviceCreate)
			items, _, _ := client.DescribeDevice(ctx, strconv.Itoa(id.DeviceID))

			// Assert
			assert.Equal(t, items.Value.ID, strconv.Itoa(id.DeviceID))
			assert.Equal(t, items.Value.Platform, tc.Platform)
			assert.Equal(t, items.Value.UserID, tc.UserId)

		})
	}
}

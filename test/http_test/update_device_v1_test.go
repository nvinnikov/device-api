//go:build httptest1
// +build httptest1

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

func TestUpdateDevice(t *testing.T) {

	var URL = envy.Get("BASE_URL", "http://localhost::8080")
	type testCasePositive struct {
		Platform    string
		UserId      string
		UpdPlatform string
		UpdUserId   string
	}
	testsUpdateDevicePositive := []testCasePositive{
		{"Ios", "111", "Android", "344"},
		{"Android", "999", "Ios", "876"},
		{"Ubuntu", "555", "Android", "321"},
	}
	for _, tc := range testsUpdateDevicePositive {
		t.Run("Update device", func(t *testing.T) {
			// Arrange
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			deviceCreate := models.CreateDeviceRequest{
				Platform: tc.Platform,
				UserID:   tc.UserId,
			}
			deviceUpdate := models.UpdateDeviceRequest{
				Platform: tc.UpdPlatform,
				UserID:   tc.UpdUserId,
			}
			ctx := context.Background()

			// Act
			id, _, err := client.CreateDevice(ctx, deviceCreate)
			assert.NoError(t, err)

			updateResult, _, err := client.UpdateDevice(ctx, strconv.Itoa(id.DeviceID), deviceUpdate)
			assert.NoError(t, err)

			// Assert
			assert.Equal(t, updateResult.Success, true, "Device updated")
		})
	}
}

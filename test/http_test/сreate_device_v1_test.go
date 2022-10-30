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

func TestCreateDevice(t *testing.T) {

	var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")
	valueLenErr := "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.Platform: value length must be at least 1 runes"
	type testCasePositive struct {
		Platform string
		UserId   string
	}
	type testCaseNegative struct {
		Platform string
		UserId   string
		err      string
	}
	testsCreateDevicePositive := []testCasePositive{
		{"Ios", "111"},
		{"Android", "999"},
		{"Ubuntu", "555"},
	}
	testCreateDeviceNegative := []testCaseNegative{
		{"", "665", valueLenErr},
	}
	for _, tc := range testsCreateDevicePositive {
		t.Run("Create device", func(t *testing.T) {
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

			// Assert
			assert.GreaterOrEqual(t, id.DeviceID, int(1))
		})
	}
	for _, tc := range testsCreateDevicePositive {
		t.Run("Create device and check description", func(t *testing.T) {

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
			description, _, err := client.DescribeDevice(ctx, strconv.Itoa(id.DeviceID))
			assert.NoError(t, err)

			// Assert
			assert.Equal(t, description.Value.ID, strconv.Itoa(id.DeviceID))
			assert.Equal(t, description.Value.Platform, tc.Platform)
			assert.Equal(t, description.Value.UserID, tc.UserId)
		})
	}
	for _, tc := range testCreateDeviceNegative {
		t.Run("Create device Negative", func(t *testing.T) {
			// Arrange
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			device := models.CreateDeviceRequest{
				Platform: tc.Platform,
				UserID:   tc.UserId,
			}
			ctx := context.Background()

			// Act
			_, resp, _ := client.CreateDevice(ctx, device)

			// Assert
			assert.Equal(t, resp.StatusCode, 400)
		})
	}
}

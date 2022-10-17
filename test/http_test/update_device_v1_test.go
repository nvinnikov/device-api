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

	var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")

	t.Run("Update device", func(t *testing.T) {
		// Arrange
		client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
		deviceCreate := models.CreateDeviceRequest{
			Platform: "Ios",
			UserID:   "700",
		}
		deviceUpdate := models.UpdateDeviceRequest{
			Platform: "Ubuntu",
			UserID:   "701",
		}
		ctx := context.Background()

		// Act
		id, _, _ := client.CreateDevice(ctx, deviceCreate)
		updateResult, _, _ := client.UpdateDevice(ctx, strconv.Itoa(id.DeviceID), deviceUpdate)

		// Assert
		assert.Equal(t, updateResult.Success, true, "Device updated")
	})

}

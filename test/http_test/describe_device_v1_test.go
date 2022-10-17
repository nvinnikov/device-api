package http_test

import (
	"context"
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"
	"time"

	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/models"

	"github.com/gobuffalo/envy"
	"github.com/stretchr/testify/assert"
)

func TestDescribeDevice(t *testing.T) {

	var URL = envy.Get("BASE_URL", "http://127.0.0.1:8080")

	t.Run("Describe device", func(t *testing.T) {
		// Arrange
		n, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err != nil {
			t.Error("error:", err)
		}
		client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
		platform, userID := "Ubuntu", strconv.Itoa(int(n.Int64()))
		deviceCreate := models.CreateDeviceRequest{
			Platform: platform,
			UserID:   userID,
		}
		ctx := context.Background()

		// Act
		id, _, _ := client.CreateDevice(ctx, deviceCreate)
		items, _, _ := client.DescribeDevice(ctx, strconv.Itoa(id.DeviceID))

		// Assert
		assert.Equal(t, items.Value.ID, strconv.Itoa(id.DeviceID))
		assert.Equal(t, items.Value.Platform, platform)
		assert.Equal(t, items.Value.UserID, userID)

	})

}

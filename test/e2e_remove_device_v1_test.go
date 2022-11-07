package test

import (
	"context"
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/stretchr/testify/assert"
	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/models"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SetupSuite struct {
	suite.Suite
}

func (s *SetupSuite) RemoveDeviceV1(t provider.T) {
	var (
		ctx context.Context
		id  string
	)
	t.Epic("Demo")
	t.Feature("E2E")
	t.Title("CreateDevice and RemoveDevice")
	t.Description(`e2e smoke case`)
	t.Tags("Smoke", "E2E", "Demo", "Ubuntu")
	t.WithTestSetup(func(t provider.T) {
		t.WithNewStep("CreateDevice", func(sCtx provider.StepCtx) {
			var URL = envy.Get("BASE_URL", "http://localhost:8080")
			client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
			device := models.CreateDeviceRequest{
				Platform: "Ubuntu",
				UserID:   "555",
			}
			ctx = context.Background()

			// Act
			id, _, err := client.CreateDevice(ctx, device)
			assert.NoError(t, err)

			// Assert
			assert.GreaterOrEqual(t, id.DeviceID, 1)
			sCtx.WithNewParameters(id.DeviceID, id)
		})
	})

	t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", id)))

	t.WithNewStep("Remove device", func(sCtx provider.StepCtx) {
		// Arrange
		var URL = envy.Get("BASE_URL", "http://localhost:8080")
		client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
		ctx := context.Background()

		// Act
		_, _, err := client.RemoveDevice(ctx, id)
		assert.NoError(t, err)

	})
	defer t.WithTestTeardown(func(t provider.T) {
		t.WithNewStep("Close ctx", func(sCtx provider.StepCtx) {
			ctx.Done()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})
}

func TestRunner(t *testing.T) {
	suite.RunSuite(t, new(SetupSuite))
}

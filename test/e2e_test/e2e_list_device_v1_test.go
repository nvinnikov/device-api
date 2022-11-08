package e2e_test

import (
	"context"
	"fmt"
	"github.com/gobuffalo/envy"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/stretchr/testify/assert"
	apiClient "gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/client"
	"gitlab.ozon.dev/qa/classroom-4/act-device-api/test/http_test/models"
	"net/url"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TestCreateAndListDeviceV1 struct {
	suite.Suite
}

func (s *TestCreateAndListDeviceV1) TestListDeviceV1(t provider.T) {
	var (
		ctx context.Context
		id  string
	)
	t.Epic("HTTP")
	t.Feature("E2E")
	t.Title("CreateDevice and ListDevice")
	t.Description(`e2e smoke case`)
	t.Tags("Smoke", "E2E", "Demo", "Ubuntu", "HTTP")
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

	t.WithNewStep("List device", func(sCtx provider.StepCtx) {
		// Arrange
		var URL = envy.Get("BASE_URL", "http://localhost:8080")
		client := apiClient.NewHTTPClient(URL, 5, 1*time.Second)
		opts := url.Values{}
		opts.Add("page", "1")
		opts.Add("perPage", "10")
		ctx := context.Background()

		// Act
		items, _, err := client.ListDevices(ctx, opts)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(items.Items), 1)

	})
	defer t.WithTestTeardown(func(t provider.T) {
		t.WithNewStep("Close ctx", func(sCtx provider.StepCtx) {
			ctx.Done()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})
}

func TestRunnerList(t *testing.T) {
	suite.RunSuite(t, new(TestCreateAndListDeviceV1))
}

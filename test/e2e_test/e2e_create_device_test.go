package e2e

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/require"
	actdeviceapi "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateDeviceAndCheckDescription(t *testing.T) {
	host := "localhost:8082"
	ctx := context.Background()
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	require.NoError(t, err)

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err: %v", err)
		}
	}(conn)

	actDeviceApiClient := actdeviceapi.NewActDeviceApiServiceClient(conn)

	runner.Run(t, "CreateDevice and check Description", func(t provider.T) {
		t.Epic("Demo")
		t.Feature("E2E")
		t.Title("CreateDevice and check Description")
		t.Description(`e2e smoke case`)
		t.Tags("Smoke", "E2E", "Demo", "Android")

		t.WithNewStep("CreateDevice", func(sCtx provider.StepCtx) {
			req := &actdeviceapi.CreateDeviceV1Request{
				Platform: "Android",
				UserId:   111,
			}
			t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req)))

			res, err := actDeviceApiClient.CreateDeviceV1(ctx, req)
			t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", res)))
			// Assert
			sCtx.Require().NoError(err)
			sCtx.Require().NotEmpty(res.DeviceId, "empty Id")
			sCtx.Require().GreaterOrEqual(res.DeviceId, uint64(1))

			t.WithNewAttachment("check Description request", allure.Text, []byte(fmt.Sprintf("%+v", res)))
			describeReq := actdeviceapi.DescribeDeviceV1Request{
				DeviceId: res.DeviceId,
			}
			description, _ := actDeviceApiClient.DescribeDeviceV1(ctx, &describeReq)
			t.WithNewAttachment("check Description response", allure.Text, []byte(fmt.Sprintf("%+v", res)))
			// Assert
			sCtx.Require().Equal(description.Value.Id, res.DeviceId)
			sCtx.Require().Equal(description.Value.Platform, "Android")
		})
	})
}

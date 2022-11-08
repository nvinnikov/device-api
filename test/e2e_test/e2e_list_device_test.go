package e2e_test

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

func TestCreateDeviceAndCheckList(t *testing.T) {
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

	runner.Run(t, "CreateDevice and check List", func(t provider.T) {
		t.Epic("GRPC")
		t.Feature("E2E")
		t.Title("CreateDevice and check list")
		t.Description(`e2e smoke case`)
		t.Tags("Smoke", "E2E", "Demo", "Ios", "Android", "GRPC")

		t.WithNewStep("CreateDevice Ios", func(sCtx provider.StepCtx) {
			req1 := &actdeviceapi.CreateDeviceV1Request{
				Platform: "Ios",
				UserId:   111,
			}
			t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req1)))

			res1, err := actDeviceApiClient.CreateDeviceV1(ctx, req1)
			t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", res1)))
			// Assert
			sCtx.Require().NoError(err)
			sCtx.Require().NotEmpty(res1.DeviceId, "empty Id")
			sCtx.Require().GreaterOrEqual(res1.DeviceId, uint64(1))
			req2 := &actdeviceapi.CreateDeviceV1Request{
				Platform: "Android",
				UserId:   111,
			}
			t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req2)))

			res2, err := actDeviceApiClient.CreateDeviceV1(ctx, req2)
			t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", res2)))
			// Assert
			sCtx.Require().NoError(err)
			sCtx.Require().NotEmpty(res2.DeviceId, "empty Id")
			sCtx.Require().GreaterOrEqual(res2.DeviceId, uint64(1))

			listReq := &actdeviceapi.ListDevicesV1Request{
				Page:    1,
				PerPage: 10,
			}
			t.WithNewAttachment("check List request", allure.Text, []byte(fmt.Sprintf("%+v", listReq)))

			list, _ := actDeviceApiClient.ListDevicesV1(ctx, listReq)
			t.WithNewAttachment("check List response", allure.Text, []byte(fmt.Sprintf("%+v", list)))
		})
	})
}

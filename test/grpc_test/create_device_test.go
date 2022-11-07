package grpc_test

import (
	"context"
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateDevice(t *testing.T) {
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

	valueLenErr := "rpc error: code = InvalidArgument desc = invalid CreateDeviceV1Request.Platform: value length must be at least 1 runes"
	type testCasePositive struct {
		Platform string
		UserId   uint64
	}
	type testCaseNegative struct {
		Platform string
		UserId   uint64
		err      string
	}
	testsCreateDevicePositive := []testCasePositive{
		{"Ios", 111},
		{"Android", 999},
		{"Ubuntu", 555},
	}
	testCreateDeviceNegative := []testCaseNegative{
		{"", 665, valueLenErr},
	}

	actDeviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)

	for _, tc := range testsCreateDevicePositive {
		runner.Run(t, "CreateDevice valid", func(t provider.T) {

			t.WithNewStep("CreateDevice", func(sCtx provider.StepCtx) {
				req := &act_device_api.CreateDeviceV1Request{
					Platform: tc.Platform,
					UserId:   tc.UserId,
				}
				t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req)))

				res, err := actDeviceApiClient.CreateDeviceV1(ctx, req)
				t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", res)))
				// Assert
				sCtx.Require().NoError(err)
				sCtx.Require().NotEmpty(res.DeviceId, "empty Id")
				sCtx.Require().GreaterOrEqual(res.DeviceId, uint64(1))
			})
		})
	}
	for _, tc := range testsCreateDevicePositive {
		runner.Run(t, "CreateDevice and check Description", func(t provider.T) {

			t.WithNewStep("CreateDevice", func(sCtx provider.StepCtx) {
				req := &act_device_api.CreateDeviceV1Request{
					Platform: tc.Platform,
					UserId:   tc.UserId,
				}
				t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req)))

				res, err := actDeviceApiClient.CreateDeviceV1(ctx, req)
				t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", res)))
				// Assert
				sCtx.Require().NoError(err)
				sCtx.Require().NotEmpty(res.DeviceId, "empty Id")
				sCtx.Require().GreaterOrEqual(res.DeviceId, uint64(1))

				t.WithNewAttachment("check Description request", allure.Text, []byte(fmt.Sprintf("%+v", res)))
				describeReq := act_device_api.DescribeDeviceV1Request{
					DeviceId: res.DeviceId,
				}
				description, _ := actDeviceApiClient.DescribeDeviceV1(ctx, &describeReq)
				t.WithNewAttachment("check Description response", allure.Text, []byte(fmt.Sprintf("%+v", res)))
				// Assert
				sCtx.Require().Equal(description.Value.Id, res.DeviceId)
				sCtx.Require().Equal(description.Value.Platform, tc.Platform)
				sCtx.Require().Equal(description.Value.UserId, tc.UserId)
			})
		})
	}
	for _, tc := range testCreateDeviceNegative {
		runner.Run(t, "CreateDevice invalid", func(t provider.T) {

			t.WithNewStep("invalid", func(sCtx provider.StepCtx) {
				req := &act_device_api.CreateDeviceV1Request{
					Platform: tc.Platform,
					UserId:   tc.UserId,
				}
				t.WithNewAttachment("CreateDevice request", allure.Text, []byte(fmt.Sprintf("%+v", req)))

				_, err := actDeviceApiClient.CreateDeviceV1(ctx, req)
				t.WithNewAttachment("CreateDevice response", allure.Text, []byte(fmt.Sprintf("%+v", err)))
				// Assert
				sCtx.Require().Equal(err.Error(), tc.err)
			})
		})
	}
}

//go:build grpctest
// +build grpctest

package grpc_test

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"

	"testing"

	"google.golang.org/grpc"
)

func TestDescribeDevice(t *testing.T) {
	host := "localhost:8082"
	ctx := context.Background()
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("grpc.Dial() err: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Logf("conn.Close err: %v", err)
		}
	}(conn)

	actDeviceApiClient := act_device_api.NewActDeviceApiServiceClient(conn)
	type testCasePositive struct {
		Platform string
		UserId   uint64
	}
	type testCaseNegative struct {
		DeviceId uint64
		result   string
	}
	testsDescribeDevicePositive := []testCasePositive{
		{"Ios", 111},
		{"Android", 999},
		{"Ubuntu", 555},
	}
	testDescribeDeviceNegative := []testCaseNegative{
		{100, "rpc error: code = NotFound desc = device not found"},
		{0, "rpc error: code = InvalidArgument desc = invalid DescribeDeviceV1Request.DeviceId: value must be greater than 0"},
		{99999, "rpc error: code = NotFound desc = device not found"},
	}
	for _, tc := range testsDescribeDevicePositive {
		t.Run("CreateDevice and check DescribeDevice", func(t *testing.T) {

			req := act_device_api.CreateDeviceV1Request{
				Platform: tc.Platform,
				UserId:   tc.UserId,
			}
			res, err := actDeviceApiClient.CreateDeviceV1(ctx, &req)
			// Assert
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.GreaterOrEqual(t, res.DeviceId, uint64(1))
			describeReq := act_device_api.DescribeDeviceV1Request{
				DeviceId: res.DeviceId,
			}
			description, _ := actDeviceApiClient.DescribeDeviceV1(ctx, &describeReq)
			// Assert
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.Equal(t, description.Value.Id, res.DeviceId)
			assert.Equal(t, description.Value.Platform, tc.Platform)
			assert.Equal(t, description.Value.UserId, tc.UserId)
			assert.Contains(t, "Android, Ios", description.Value.Platform)
			assert.NotEmpty(t, description.Value.UserId)
			assert.NotEmpty(t, description.Value.EnteredAt)
		})
	}
	for _, tc := range testDescribeDeviceNegative {
		t.Run("DescribeDevice not existing", func(t *testing.T) {
			req := act_device_api.DescribeDeviceV1Request{
				DeviceId: tc.DeviceId,
			}

			_, err := actDeviceApiClient.DescribeDeviceV1(ctx, &req)
			assert.Equal(t, err.Error(), tc.result)
		})
	}
}

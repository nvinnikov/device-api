//go:build grpctest
// +build grpctest

package grpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
)

func TestUpdateDevice(t *testing.T) {
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
		Platform    string
		UserId      uint64
		UpdPlatform string
		UpdUserId   uint64
	}
	testsUpdateDevicePositive := []testCasePositive{
		{"Ios", 111, "Android", 344},
		{"Android", 999, "Ios", 876},
		{"Ubuntu", 555, "Android", 321},
	}
	for _, tc := range testsUpdateDevicePositive {
		t.Run("UpdateDevice valid", func(t *testing.T) {
			req := act_device_api.CreateDeviceV1Request{
				Platform: tc.Platform,
				UserId:   tc.UserId,
			}

			res, err := actDeviceApiClient.CreateDeviceV1(ctx, &req)
			require.NoError(t, err)
			require.NotNil(t, res)

			assert.GreaterOrEqual(t, res.DeviceId, uint64(1))

			reqUpd := act_device_api.UpdateDeviceV1Request{
				DeviceId: res.DeviceId,
				Platform: tc.UpdPlatform,
				UserId:   tc.UpdUserId,
			}
			resUpd, err := actDeviceApiClient.UpdateDeviceV1(ctx, &reqUpd)
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.Equal(t, resUpd.Success, true, "Device updated")
			describeReq := act_device_api.DescribeDeviceV1Request{
				DeviceId: res.DeviceId,
			}
			description, _ := actDeviceApiClient.DescribeDeviceV1(ctx, &describeReq)
			// Assert
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.Equal(t, description.Value.Id, res.DeviceId)
			assert.Equal(t, description.Value.Platform, tc.UpdPlatform)
			assert.Equal(t, description.Value.UserId, tc.UpdUserId)
			assert.Contains(t, "Android, Ios", description.Value.Platform)
			assert.NotEmpty(t, description.Value.UserId)
			assert.NotEmpty(t, description.Value.EnteredAt)

		})
	}
}

//go:build grpctest
// +build grpctest

package grpc_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateDevice(t *testing.T) {
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
		t.Run("CreateDevice valid", func(t *testing.T) {
			req := act_device_api.CreateDeviceV1Request{
				Platform: tc.Platform,
				UserId:   tc.UserId,
			}
			res, err := actDeviceApiClient.CreateDeviceV1(ctx, &req)

			// Assert
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.GreaterOrEqual(t, res.DeviceId, uint64(1))

		})
	}
	for _, tc := range testsCreateDevicePositive {
		t.Run("CreateDevice and check Description", func(t *testing.T) {
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
			assert.Equal(t, description.Value.Id, res.DeviceId)
			assert.Equal(t, description.Value.Platform, tc.Platform)
			assert.Equal(t, description.Value.UserId, tc.UserId)

		})
	}
	for _, tc := range testCreateDeviceNegative {
		t.Run("CreateDevice invalid", func(t *testing.T) {
			req := act_device_api.CreateDeviceV1Request{
				Platform: tc.Platform,
				UserId:   tc.UserId,
			}
			_, err := actDeviceApiClient.CreateDeviceV1(ctx, &req)
			assert.Equal(t, err.Error(), tc.err)
		})
	}
}

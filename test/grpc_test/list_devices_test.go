//go:build grpctest
// +build grpctest

package grpc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"

	"testing"

	"google.golang.org/grpc"
)

func TestListDevices(t *testing.T) {
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
		Page    uint64
		PerPage uint64
	}
	testsListDevicesPositive := []testCasePositive{
		{1, 10},
		{1, 20},
		{2, 10},
	}
	testsListDevicesNegative := []testCasePositive{
		{0, 10},
	}
	for _, tc := range testsListDevicesPositive {
		t.Run("DescribeDevice existing", func(t *testing.T) {
			req := act_device_api.ListDevicesV1Request{
				Page:    tc.Page,
				PerPage: tc.PerPage,
			}

			res, err := actDeviceApiClient.ListDevicesV1(ctx, &req)
			require.NoError(t, err)
			require.NotNil(t, res)
		})
	}
	for _, tc := range testsListDevicesNegative {
		t.Run("DescribeDevice negative", func(t *testing.T) {
			req := act_device_api.ListDevicesV1Request{
				Page:    tc.Page,
				PerPage: tc.PerPage,
			}

			_, err := actDeviceApiClient.ListDevicesV1(ctx, &req)
			require.NotNil(t, err)
		})
	}
}

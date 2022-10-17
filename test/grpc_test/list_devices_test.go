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

	t.Run("DescribeDevice existing", func(t *testing.T) {
		req := act_device_api.ListDevicesV1Request{
			Page:    1,
			PerPage: 10,
		}

		res, err := actDeviceApiClient.ListDevicesV1(ctx, &req)
		require.NoError(t, err)
		require.NotNil(t, res)

	})

}

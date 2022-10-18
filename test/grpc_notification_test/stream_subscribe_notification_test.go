package grpc_notification

import (
	"context"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestSubscribeNotification(t *testing.T) {
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

	actDeviceApiClient := act_device_api.NewActNotificationApiServiceClient(conn)

	t.Run("GetNotification valid", func(t *testing.T) {
		req := act_device_api.SubscribeNotificationRequest{
			DeviceId: 1,
		}

		res, err := actDeviceApiClient.SubscribeNotification(ctx, &req)
		require.NoError(t, err)
		require.NotNil(t, res)

	})

}

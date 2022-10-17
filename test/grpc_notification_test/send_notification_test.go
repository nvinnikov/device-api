package grpc_notification_test

import (
	"context"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
	"google.golang.org/grpc"
	"testing"
)

func TestSendNotification(t *testing.T) {
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

	t.Run("SendNotification valid", func(t *testing.T) {

		notificationReq := act_device_api.Notification{
			NotificationId:     1,
			DeviceId:           1,
			Username:           "Nikita",
			Message:            "Hello",
			Lang:               0,
			NotificationStatus: 0,
		}
		req := act_device_api.SendNotificationV1Request{
			Notification: &notificationReq,
		}

		res, err := actDeviceApiClient.SendNotificationV1(ctx, &req)
		require.NoError(t, err)
		require.NotNil(t, res)
	})
}

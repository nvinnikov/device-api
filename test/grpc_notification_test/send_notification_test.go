package grpc_notification_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
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
	type testCase struct {
		NotificationId     uint64
		DeviceId           uint64
		Username           string
		Message            string
		Lang               act_device_api.Language
		NotificationStatus act_device_api.Status
	}

	testsSendNotification := []testCase{
		{111, 1, "nikita", "nikita", act_device_api.Language_LANG_ENGLISH, act_device_api.Status_STATUS_CREATED},
		{112, 1, "nikita", "nikita", act_device_api.Language_LANG_RUSSIAN, act_device_api.Status_STATUS_CREATED},
		{113, 1, "nikita", "nikita", act_device_api.Language_LANG_ESPANOL, act_device_api.Status_STATUS_CREATED},
		{114, 1, "nikita", "nikita", act_device_api.Language_LANG_ITALIAN, act_device_api.Status_STATUS_CREATED},
	}
	actDeviceApiClient := act_device_api.NewActNotificationApiServiceClient(conn)
	for _, tc := range testsSendNotification {

		t.Run("SendNotification valid", func(t *testing.T) {

			notificationReq := act_device_api.Notification{
				NotificationId:     tc.NotificationId,
				DeviceId:           tc.DeviceId,
				Username:           tc.Username,
				Message:            tc.Message,
				Lang:               tc.Lang,
				NotificationStatus: tc.NotificationStatus,
			}
			req := act_device_api.SendNotificationV1Request{
				Notification: &notificationReq,
			}

			res, err := actDeviceApiClient.SendNotificationV1(ctx, &req)
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.GreaterOrEqual(t, res.NotificationId, uint64(1))

		})
	}
}

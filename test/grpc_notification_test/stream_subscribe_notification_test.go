//go:build grpctest_n
// +build grpctest_n

package grpc_notification

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"google.golang.org/grpc"

	act_device_api "gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api/gitlab.ozon.dev/qa/classroom-4/act-device-api/pkg/act-device-api"
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

		res, _ := actDeviceApiClient.SubscribeNotification(ctx, &req)

		host := "localhost:8082"
		ctx := context.Background()
		conn, _ := grpc.Dial(host, grpc.WithInsecure())
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
			result             string
		}

		testsSendNotification := []testCase{
			{111, 1, "nikita", "nikita", act_device_api.Language_LANG_ENGLISH, act_device_api.Status_STATUS_CREATED, "Good afternoon nikita"},
			{112, 1, "nikita", "nikita", act_device_api.Language_LANG_RUSSIAN, act_device_api.Status_STATUS_CREATED, "Добрый вечер nikita"},
			{113, 1, "nikita", "nikita", act_device_api.Language_LANG_ESPANOL, act_device_api.Status_STATUS_CREATED, "Buenas noches nikita"},
			{114, 1, "nikita", "nikita", act_device_api.Language_LANG_ITALIAN, act_device_api.Status_STATUS_CREATED, "Buona serata nikita"},
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

				notificationRes, err := actDeviceApiClient.SendNotificationV1(ctx, &req)
				require.NoError(t, err)
				require.NotNil(t, notificationRes)
				assert.GreaterOrEqual(t, notificationRes.NotificationId, uint64(1))
				msg, err := res.Recv()
				if err != nil {
					t.Fatalf("error while reading stream: %v", err)
				}
				assert.Equal(t, tc.result, msg.Message)
				fmt.Printf("Response from SendNotification: %v %v\n", msg.NotificationId, msg.Message)
			})
		}
		//for {
		//	msg, err := res.Recv()
		//	if err == io.EOF {
		//		break
		//	}
		//	if err != nil {
		//		t.Fatalf("error while reading stream: %v", err)
		//	}
		//	fmt.Printf("Response from SendNotification: %v %v\n", msg.NotificationId, msg.Message)
		//}

	})

}

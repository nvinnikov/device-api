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

	t.Run("UpdateDevice valid", func(t *testing.T) {
		req := act_device_api.CreateDeviceV1Request{
			Platform: "Android",
			UserId:   123000,
		}

		res, err := actDeviceApiClient.CreateDeviceV1(ctx, &req)
		require.NoError(t, err)
		require.NotNil(t, res)

		assert.GreaterOrEqual(t, res.DeviceId, uint64(1))

		reqUpd := act_device_api.UpdateDeviceV1Request{
			DeviceId: res.DeviceId,
			Platform: "Ios",
			UserId:   666,
		}
		resUpd, err := actDeviceApiClient.UpdateDeviceV1(ctx, &reqUpd)
		require.NoError(t, err)
		require.NotNil(t, res)

		assert.Equal(t, resUpd.Success, true, "Device updated")

	})

}

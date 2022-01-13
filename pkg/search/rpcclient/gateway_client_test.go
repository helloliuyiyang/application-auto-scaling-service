package rpcclient

import (
	"net"
	"testing"

	"nanto.io/application-auto-scaling-service/pkg/common/config"
	"nanto.io/application-auto-scaling-service/pkg/search/types"
)

func TestAppGatewayClient_AddServer(t *testing.T) {
	if err := InitAppGwReporter(&config.AppGatewayConf{
		GrpcHost: "127.0.0.1",
		GrpcPort: 9999,
	}); err != nil {
		t.Fatal(err)
	}
	if err := GetAppGwReporter().AddServer("fleetId", types.Server{
		Id:            "xxxxxxxxxxxxxxx",
		IP:            net.ParseIP("127.0.0.1"),
		MaxProcessNum: 10,
		ProcessConfig: &types.ProcessConfig{MaxConcurrency: 10},
	}); err != nil {
		t.Error(err)
	}
}

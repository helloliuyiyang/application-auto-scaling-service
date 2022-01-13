package rpcclient

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"nanto.io/application-auto-scaling-service/pkg/common/config"
	"nanto.io/application-auto-scaling-service/pkg/common/utils/logutil"
	apis "nanto.io/application-auto-scaling-service/pkg/search/rpcclient/pbfiles"
	"nanto.io/application-auto-scaling-service/pkg/search/types"
)

var logger = logutil.GetLogger()

var gwReporter *AppGatewayReporter

type AppGatewayReporter struct {
	apis.MataServiceClient
	conn *grpc.ClientConn
}

func GetAppGwReporter() *AppGatewayReporter {
	if gwReporter == nil {
		logger.Fatal("AppGatewayReporter must be initialized")
	}
	return gwReporter
}

func InitAppGwReporter(conf *config.AppGatewayConf) error {
	address := fmt.Sprintf("%s:%d", conf.GrpcHost, conf.GrpcPort)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return errors.Wrapf(err, "Grpc dial address[%s] err", address)
	}

	logger.Infof("Init application-gateway[addr: %s] client success", address)
	gwReporter = &AppGatewayReporter{
		MataServiceClient: apis.NewMataServiceClient(conn),
		conn:              conn,
	}
	return nil
}

func (c *AppGatewayReporter) Close() {
	if err := c.conn.Close(); err != nil {
		logger.Errorf("Close gateway grpc conn err: %v", err)
	}
}

// AddFleet 创建AS
func (c *AppGatewayReporter) AddFleet(fleetId, fleetName string) error {
	// todo grpc超时控制
	req := &apis.CreateServerQueueReq{
		Name: fleetName,
	}
	resp, err := c.CreateServerQueue(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "Grpc CreateServerQueue err")
	}

	if resp.GetErrMsg().GetStatusCode() != apis.StatusCode_CreateServerQueueSuccess {
		return errors.Errorf("Grpc CreateServerQueue err, code: %d", resp.GetErrMsg().GetStatusCode())
	}
	return nil
}

// AddServer 给fleet添加server实例
func (c *AppGatewayReporter) AddServer(fleetId string, server types.Server) error {
	// todo grpc超时控制
	req := &apis.CreateServerReq{
		ServerQueueID: fleetId,
		Server: &apis.Server{
			Id:            server.Id,
			Ip:            server.IP.String(),
			ProcessConfig: &apis.ProcessConfig{MaxConcurrency: server.ProcessConfig.MaxConcurrency},
		},
	}
	resp, err := c.CreateServer(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "Grpc CreateServerQueue err")
	}

	if resp.GetErrMsg().GetStatusCode() != apis.StatusCode_CreateServerSuccess {
		return errors.Errorf("Grpc CreateServerQueue err, code: %d", resp.GetErrMsg().GetStatusCode())
	}
	logger.Infof("Grpc create server[%s] success", server.Id)
	return nil
}

package grpc_client

import (
	"context"
	"github.com/brocaar/chirpstack-api/go/v3/as/external/api"
)

func (g *GrpcClient) DeviceEnqueue(ctx context.Context, serverParam ServerParam, devEUI string, FPort uint32, confirmed bool, data []byte) (FCnt uint32, err error) {
	// init gRPC dial options
	conn, err := g.initConn(serverParam)
	if err != nil {
		return 0, err
	}
	client := api.NewDeviceQueueServiceClient(conn)
	// make an Enqueue api call
	resp, err := client.Enqueue(ctx, &api.EnqueueDeviceQueueItemRequest{
		DeviceQueueItem: &api.DeviceQueueItem{
			DevEui:    devEUI,
			FPort:     FPort,
			Confirmed: confirmed,
			Data:      data,
		},
	})
	if err != nil {
		return 0, err
	}
	return resp.FCnt, err
}
func (g *GrpcClient) Login(ctx context.Context, serverUrl, email, password string) (string, error) {
	// init gRPC dial options
	conn, err := g.initConn(ServerParam{
		ServerUrl: serverUrl,
	})
	if err != nil {
		return "", err
	}
	client := api.NewInternalServiceClient(conn)
	resp, err := client.Login(ctx, &api.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return resp.Jwt, nil
}
func (g *GrpcClient) GetOrgList(ctx context.Context, serverParam ServerParam) ([]*api.OrganizationListItem, error) {
	// init gRPC dial options
	conn, err := g.initConn(serverParam)
	if err != nil {
		return nil, err
	}
	resp, err := api.NewOrganizationServiceClient(conn).List(ctx, &api.ListOrganizationRequest{
		Offset: 0,
		Limit:  10000,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}
func (g *GrpcClient) GetApplicationList(ctx context.Context, serverParam ServerParam, orgID int64) ([]*api.ApplicationListItem, error) {
	// init gRPC dial options
	conn, err := g.initConn(serverParam)
	if err != nil {
		return nil, err
	}
	resp, err := api.NewApplicationServiceClient(conn).List(ctx, &api.ListApplicationRequest{
		OrganizationId: orgID,
		Limit:          10000,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

func (g *GrpcClient) GetDeviceList(ctx context.Context, serverParam ServerParam, appID int64) ([]*api.DeviceListItem, error) {
	// init gRPC dial options
	conn, err := g.initConn(serverParam)
	if err != nil {
		return nil, err
	}
	resp, err := api.NewDeviceServiceClient(conn).List(ctx, &api.ListDeviceRequest{
		ApplicationId: appID,
		Limit:         10000,
	})
	if err != nil {
		return nil, err
	}
	return resp.Result, nil
}

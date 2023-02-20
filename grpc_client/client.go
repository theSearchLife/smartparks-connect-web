package grpc_client

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

var DefaultGrpcClient *GrpcClient

func init() {
	DefaultGrpcClient = &GrpcClient{
		grpcConns: make(map[ServerParam]*grpc.ClientConn),
		lock:      sync.RWMutex{},
	}
}

type ServerParam struct {
	ServerUrl string
	ApiToken  string
}

type GrpcClient struct {
	grpcConns map[ServerParam]*grpc.ClientConn
	lock      sync.RWMutex
}

func (g *GrpcClient) initConn(serverParam ServerParam) (*grpc.ClientConn, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	if client, ok := g.grpcConns[serverParam]; ok {
		return client, nil
	}
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(serverParam.ApiToken)),
		grpc.WithTimeout(time.Second * 3),
		grpc.WithInsecure(),
	}

	// connect to the gRPC server
	conn, err := grpc.Dial(serverParam.ServerUrl, dialOpts...)
	if err != nil {
		return nil, err
	}
	g.grpcConns[serverParam] = conn
	return conn, err
}

type APIToken string

func (a APIToken) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a),
	}, nil
}

func (a APIToken) RequireTransportSecurity() bool {
	return false
}

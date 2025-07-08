package shared

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/hdget/lib-jobexecutor/protobuf"
	"google.golang.org/grpc"
)

type TaskPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl protobuf.TaskServer
}

func (t TaskPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	protobuf.RegisterTaskServer(server, t.Impl)
	return nil
}

func (t TaskPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	return &TaskClient{
		client: protobuf.NewTaskClient(conn),
	}, nil
}

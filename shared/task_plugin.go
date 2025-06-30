package shared

import (
	"context"
	"github.com/hashicorp/go-plugin"
	"github.com/hdget/lib-jobexecutor/protobuf"
	"google.golang.org/grpc"
)

// GetHandshakeConfig 确保主程序与插件版本兼容
func GetHandshakeConfig(token string) plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "TASK_PLUGIN",
		MagicCookieValue: token,
	}
}

type TaskPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl protobuf.TaskServer
	Job  protobuf.JobClient
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

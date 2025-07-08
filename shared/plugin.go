package shared

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/hdget/lib-jobexecutor/protobuf"
	"log"
	"os/exec"
)

const (
	envJobHost = "HD_JOB_ADDRESS"
)

var (
	pluginHandshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "TASK_PLUGIN",
		MagicCookieValue: "sK2jX8",
	}
)

func NewPluginClient(jobServerAddress, taskPluginPath, taskPluginName string, logger *log.Logger) *plugin.Client {
	// 将宿主服务的地址注入到环境变量中传递给插件
	cmd := exec.Command(taskPluginPath)
	cmd.Env = []string{fmt.Sprintf("%s=%s", envJobHost, jobServerAddress)}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  pluginHandshakeConfig,
		Plugins:          map[string]plugin.Plugin{taskPluginName: &TaskPlugin{}},
		Cmd:              cmd,
		SyncStderr:       logger.Writer(),
		SyncStdout:       logger.Writer(),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})

	return client
}

func PluginServe(taskName string, taskServer protobuf.TaskServer) {
	// 2. start plugin serve
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: pluginHandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			taskName: &TaskPlugin{
				Impl: taskServer,
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

package shared

import (
	"context"
	"github.com/hdget/lib-jobexecutor/protobuf"
	"github.com/pkg/errors"
)

type Task interface {
	Execute(ctx context.Context, request *protobuf.TaskExecuteRequest) (*protobuf.TaskExecuteResponse, error)
}

// TaskClient 实现 Task 接口的客户端适配器
type TaskClient struct {
	client protobuf.TaskClient
}

func (c *TaskClient) Execute(params map[string]string) (string, error) {
	resp, err := c.client.Execute(context.Background(), &protobuf.TaskExecuteRequest{
		Params: params,
	})
	if err != nil {
		return "", err
	}

	if resp.Error != "" {
		return "", errors.New(resp.Error)
	}

	return resp.Result, nil
}

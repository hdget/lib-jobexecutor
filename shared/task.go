package shared

import (
	"context"
	"github.com/hdget/lib-jobexecutor/protobuf"
)

type Task interface {
	Execute(ctx context.Context, request *protobuf.ExecuteTaskRequest) (*protobuf.ExecuteTaskResponse, error) // 执行任务
	GetDescription(ctx context.Context) (string, error)                                                       // 获取任务描述
}

// TaskClient 实现Task接口的客户端适配器
type TaskClient struct {
	client protobuf.TaskClient
}

func (c *TaskClient) Execute(request []byte) ([]byte, error) {
	result, err := c.client.Execute(context.Background(), &protobuf.ExecuteTaskRequest{
		Request: request,
	})
	if err != nil {
		return nil, err
	}
	return result.Response, nil
}

func (c *TaskClient) GetDescription() (string, error) {
	result, err := c.client.GetDescription(context.Background(), &protobuf.GetTaskDescriptionRequest{})
	if err != nil {
		return "", err
	}
	return result.Description, nil
}

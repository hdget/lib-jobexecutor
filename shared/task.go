package shared

import (
	"context"
	"github.com/hdget/lib-jobexecutor/protobuf"
)

type Task interface {
	Execute(ctx context.Context, jobId int64, jobRequest []byte) (*protobuf.ExecuteTaskResponse, error) // 执行任务
	GetDescription() (string, error)                                                                    // 获取任务描述
}

// TaskClient 实现Task接口的客户端适配器
type TaskClient struct {
	client protobuf.TaskClient
}

func (c *TaskClient) Execute(ctx context.Context, jobId int64, jobRequest []byte) (*protobuf.ExecuteTaskResponse, error) {
	response, err := c.client.Execute(ctx, &protobuf.ExecuteTaskRequest{
		JobId:      jobId,
		JobRequest: jobRequest,
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *TaskClient) GetDescription() (string, error) {
	result, err := c.client.GetDescription(context.Background(), &protobuf.GetTaskDescriptionRequest{})
	if err != nil {
		return "", err
	}
	return result.Description, nil
}

package shared

import "github.com/hdget/lib-jobexecutor/protobuf"

type Job interface {
	Invoke(request *protobuf.JobInvokeRequest) (*protobuf.JobInvokeResponse, error)
}

package shared

import "github.com/hdget/lib-jobexecutor/protobuf"

type Job interface {
	InvokeService(request *protobuf.JobInvokeServiceRequest) (*protobuf.JobInvokeServiceResponse, error)
}

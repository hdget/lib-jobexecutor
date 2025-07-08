package shared

import (
	"fmt"
	"github.com/hdget/lib-jobexecutor/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type Job interface {
	InvokeService(request *protobuf.JobInvokeServiceRequest) (*protobuf.JobInvokeServiceResponse, error)
}

func NewJobClient() (protobuf.JobClient, error) {
	// connect to job rpc server
	jobServerAddress, found := os.LookupEnv(envJobHost)
	if !found {
		return nil, fmt.Errorf("%s environment variable not set", envJobHost)
	}

	conn, err := grpc.NewClient(jobServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed connect to job server: %s\n", jobServerAddress)
	}

	return protobuf.NewJobClient(conn), nil
}

package cloud

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type cloudServer struct {
}

func (c *cloudServer) CreateMachines(ctx context.Context, r *CreateMachinesRequest) (*CreateMachinesResponse, error) {

	if r == nil {
		err := status.Error(codes.InvalidArgument, "no request provided")
		log.Printf("Error while creating machines, error: %s\n")
		return nil, err
	}

	log.Printf("Creating machines with name [%s] with provider: [%d]. Kind: [%s], Min: [%d], Max: [%d]\n",
		r.GetName(), r.GetProvider(), r.GetKind(), r.GetMin(), r.GetMax())

	amount := r.GetMax() + r.GetMin()/2

	var i uint64
	instances := make([]string, amount)
	for i = 0; i < amount; i++ {
		instances[i] = fmt.Sprintf("%s-%d", r.GetName(), i+1)
	}

	return &CreateMachinesResponse{
		Instances: instances,
		Amount:    amount,
		Provider:  r.GetProvider(),
	}, nil
}

func (c *cloudServer) mustEmbedUnimplementedCloudServer() {
	panic("implement me")
}

func NewCloudServer() CloudServer {
	return &cloudServer{}
}

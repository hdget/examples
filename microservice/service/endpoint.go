package service

import (
	"context"
	"examples/microservice/autogen"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type Endpoints struct {
	SearchEndpoint kitgrpc.Handler
	HelloEndpoint kitgrpc.Handler
}

func (e Endpoints) Search(ctx context.Context, request *autogen.SearchRequest) (*autogen.SearchResponse, error) {
	_, resp, err := e.SearchEndpoint.ServeGRPC(ctx, request)
	return resp.(*autogen.SearchResponse), err
}

func (e Endpoints) Hello(ctx context.Context, request *autogen.HelloRequest) (*autogen.HelloResponse, error) {
	_, resp, err := e.HelloEndpoint.ServeGRPC(ctx, request)
	return resp.(*autogen.HelloResponse), err
}

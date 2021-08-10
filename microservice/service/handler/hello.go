package handler

import (
	"context"
	"examples/microservice/autogen"
	"examples/microservice/service"
	"github.com/go-kit/kit/endpoint"
)

type HelloHandler struct {}
func (h HelloHandler) GetName() string {
	return "hello"
}
func (h HelloHandler) MakeEndpoint(svc interface{}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.(*service.SearchServiceImpl).Hello(ctx, request.(*autogen.HelloRequest))
	}
}
func (h HelloHandler) ServerDecodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	panic("implement me")
}
func (h HelloHandler) ServerEncodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	panic("implement me")
}

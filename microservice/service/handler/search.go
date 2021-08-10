package handler

import (
	"context"
	"examples/microservice/autogen"
	"examples/microservice/service"
	"github.com/go-kit/kit/endpoint"
)

// put into handler file
type SearchHandler struct {}

func (s SearchHandler) GetName() string {
	return "search"
}

func (s SearchHandler) MakeEndpoint(svc interface{}) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.(*service.SearchServiceImpl).Search(ctx, request.(*autogen.SearchRequest))
	}
}

func (s SearchHandler) ServerDecodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*autogen.SearchRequest), nil
}

func (s SearchHandler) ServerEncodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	return response.(*autogen.SearchResponse), nil
}

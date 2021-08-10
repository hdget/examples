package service

import (
	"context"
	"examples/microservice/autogen"
)

type SearchServiceImpl struct {}

func (s SearchServiceImpl) Hello(ctx context.Context, request *autogen.HelloRequest) (*autogen.HelloResponse, error) {
	return &autogen.HelloResponse{
		Response: "hello world",
	}, nil
}
func (s SearchServiceImpl) Search(ctx context.Context, request *autogen.SearchRequest) (*autogen.SearchResponse, error) {
	return &autogen.SearchResponse{
		Response: "search response",
	}, nil
}

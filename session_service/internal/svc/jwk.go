package svc

import (
	"context"
	pb "github.com/vgeorgo/grpc-envoy-experiment/session_service_go/pkg/pb/session/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type JWTServer struct{}

func (s *JWTServer) Check(ctx context.Context, req *pb.CheckRequest) (res *pb.Response, err error) {
	res = &pb.Response{
		Id: "user_session_token",
	}

	return
}

func (s *JWTServer) Create(ctx context.Context, req *pb.CreateRequest) (res *pb.Response, err error) {
	if len(req.User) == 0 {
		err = status.New(codes.InvalidArgument, "User cannot be empty").Err()
		return
	}

	// Generate token

	res = &pb.Response{
		Id: "user_session_token",
	}

	return
}

func (s *JWTServer) Remove(ctx context.Context, req *pb.RemoveRequest) (res *pb.Response, err error) {
	res = &pb.Response{
		Id: "user_session_token",
	}

	return
}

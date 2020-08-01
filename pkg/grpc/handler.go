package grpc

import (
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
	endpoint "icart/pkg/endpoint"
	pb "icart/pkg/grpc/pb"
)

// makeSignInHandler creates the handler logic
func makeSignInHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SignInEndpoint, decodeSignInRequest, encodeSignInResponse, options...)
}

// decodeSignInResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SignIn request.
// TODO implement the decoder
func decodeSignInRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Icart' Decoder is not impelemented")
}

// encodeSignInResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSignInResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Icart' Encoder is not impelemented")
}
func (g *grpcServer) SignIn(ctx context1.Context, req *pb.SignInRequest) (*pb.SignInReply, error) {
	_, rep, err := g.signIn.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SignInReply), nil
}

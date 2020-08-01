package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "icart/pkg/service"
)

// SignInRequest collects the request parameters for the SignIn method.
type SignInRequest struct {
	Username string `json:"username"`
	Psw      string `json:"psw"`
}

// SignInResponse collects the response parameters for the SignIn method.
type SignInResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSignInEndpoint returns an endpoint that invokes SignIn on the service.
func MakeSignInEndpoint(s service.IcartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignInRequest)
		rs, err := s.SignIn(ctx, req.Username, req.Psw)
		return SignInResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SignInResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SignIn implements Service. Primarily useful in a client.
func (e Endpoints) SignIn(ctx context.Context, username string, psw string) (rs string, err error) {
	request := SignInRequest{
		Psw:      psw,
		Username: username,
	}
	response, err := e.SignInEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignInResponse).Rs, response.(SignInResponse).Err
}

package service

import "context"

// IcartService describes the service.
type IcartService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SignIn(ctx context.Context, username string, psw string) (rs string, err error)
}

type basicIcartService struct{}

func (b *basicIcartService) SignIn(ctx context.Context, username string, psw string) (rs string, err error) {
	// TODO implement the business logic of SignIn
	return rs, err
}

// NewBasicIcartService returns a naive, stateless implementation of IcartService.
func NewBasicIcartService() IcartService {
	return &basicIcartService{}
}

// New returns a IcartService with all of the expected middleware wired in.
func New(middleware []Middleware) IcartService {
	var svc IcartService = NewBasicIcartService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(IcartService) IcartService

type loggingMiddleware struct {
	logger log.Logger
	next   IcartService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a IcartService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next IcartService) IcartService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SignIn(ctx context.Context, username string, psw string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "SignIn", "username", username, "psw", psw, "rs", rs, "err", err)
	}()
	return l.next.SignIn(ctx, username, psw)
}

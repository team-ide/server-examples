package server_service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/team-ide/server-examples/go/gen-go/vrv_service"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(vrv_service.Service1) vrv_service.Service1

// LoggingMiddleware takes a logger as a dependency
// and returns a service Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next vrv_service.Service1) vrv_service.Service1 {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   vrv_service.Service1
}

func (mw loggingMiddleware) Hello(ctx context.Context, request *vrv_service.Service1HelloReuqest) (response *vrv_service.Service1HelloResponse, err error) {
	defer func() {
		mw.logger.Log("method", "Hello", "request", request, "response", response, "err", err)
	}()
	response, err = mw.next.Hello(ctx, request)
	return
}

// InstrumentingMiddleware returns a service middleware that instruments
// the number of integers summed and characters concatenated over the lifetime of
// the service.
func InstrumentingMiddleware(ints, chars metrics.Counter) Middleware {
	return func(next vrv_service.Service1) vrv_service.Service1 {
		return instrumentingMiddleware{
			ints:  ints,
			chars: chars,
			next:  next,
		}
	}
}

type instrumentingMiddleware struct {
	ints  metrics.Counter
	chars metrics.Counter
	next  vrv_service.Service1
}

func (mw instrumentingMiddleware) Hello(ctx context.Context, request *vrv_service.Service1HelloReuqest) (response *vrv_service.Service1HelloResponse, err error) {
	v, err := mw.next.Hello(ctx, request)
	// mw.ints.Add(float64(v))
	return v, err
}

package server_endpoint

import (
	"context"

	"golang.org/x/time/rate"

	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"

	"github.com/team-ide/server-examples/go/gen-go/vrv_service"
)

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	HelloEndpoint endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc vrv_service.Service1, logger log.Logger, duration metrics.Histogram, otTracer stdopentracing.Tracer) Set {

	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = MakeHellowEndpoint(svc)
		// Concat is limited to 1 request per second with burst of 100 requests.
		// Note, rate is defined as a number of requests per second.
		helloEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Limit(1), 100))(helloEndpoint)
		helloEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(helloEndpoint)
		helloEndpoint = opentracing.TraceServer(otTracer, "Hello")(helloEndpoint)
		helloEndpoint = LoggingMiddleware(log.With(logger, "method", "Hello"))(helloEndpoint)
		helloEndpoint = InstrumentingMiddleware(duration.With("method", "Hello"))(helloEndpoint)
	}
	return Set{
		HelloEndpoint: helloEndpoint,
	}
}

func (s Set) Hello(ctx context.Context, request *vrv_service.Service1HelloReuqest) (*vrv_service.Service1HelloResponse, error) {
	resp, err := s.HelloEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	response := resp.(*vrv_service.Service1HelloResponse)
	return response, nil
}

func MakeHellowEndpoint(s vrv_service.Service1) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*vrv_service.Service1HelloReuqest)
		return s.Hello(ctx, req)
	}
}

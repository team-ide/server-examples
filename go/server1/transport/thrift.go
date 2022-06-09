package server_transport

import (
	"context"
	"time"

	"golang.org/x/time/rate"

	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"

	"github.com/team-ide/server-examples/go/gen-go/vrv_service"
	server_endpoint "github.com/team-ide/server-examples/go/server1/endpoint"
	server_service "github.com/team-ide/server-examples/go/server1/service"
)

type thriftServer struct {
	ctx       context.Context
	endpoints server_endpoint.Set
}

// NewThriftServer makes a set of endpoints available as a Thrift service.
func NewThriftServer(endpoints server_endpoint.Set) vrv_service.Service1 {
	return &thriftServer{
		endpoints: endpoints,
	}
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func (s *thriftServer) Hello(ctx context.Context, r *vrv_service.Service1HelloReuqest) (*vrv_service.Service1HelloResponse, error) {
	response, err := s.endpoints.HelloEndpoint(ctx, r)
	if err != nil {
		return nil, err
	}
	resp := response.(*vrv_service.Service1HelloResponse)
	return resp, nil
}

// NewThriftClient returns an AddService backed by a Thrift server described by
// the provided client. The caller is responsible for constructing the client,
// and eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewThriftClient(client *vrv_service.Service1Client) vrv_service.Service1 {
	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	// Each individual endpoint is an http/transport.Client (which implements
	// endpoint.Endpoint) that gets wrapped with various middlewares. If you
	// could rely on a consistent set of client behavior.
	var helloEndpoint endpoint.Endpoint
	{
		helloEndpoint = MakeThriftHelloEndpoint(client)
		helloEndpoint = limiter(helloEndpoint)
		helloEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Hello",
			Timeout: 30 * time.Second,
		}))(helloEndpoint)
	}

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	return server_endpoint.Set{
		HelloEndpoint: helloEndpoint,
	}
}

func MakeThriftHelloEndpoint(client *vrv_service.Service1Client) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*vrv_service.Service1HelloReuqest)
		reply, err := client.Hello(ctx, req)
		if err == server_service.ErrIntOverflow {
			return nil, err // special case; see comment on ErrIntOverflow
		}
		return reply, nil
	}
}

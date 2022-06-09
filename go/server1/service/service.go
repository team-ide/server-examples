package server_service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/team-ide/server-examples/go/gen-go/vrv_service"
)

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger, ints, chars metrics.Counter) vrv_service.Service1 {
	var svc vrv_service.Service1
	{
		svc = NewBasicService()
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware(ints, chars)(svc)
	}
	return svc
}

var (
	// ErrTwoZeroes is an arbitrary business rule for the Add method.
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	// ErrIntOverflow protects the Add method. We've decided that this error
	// indicates a misbehaving service and should count against e.g. circuit
	// breakers. So, we return it directly in endpoints, to illustrate the
	// difference. In a real service, this probably wouldn't be the case.
	ErrIntOverflow = errors.New("integer overflow")

	// ErrMaxSizeExceeded protects the Concat method.
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService() vrv_service.Service1 {
	return basicService{}
}

type basicService struct{}

func (s basicService) Hello(ctx context.Context, request *vrv_service.Service1HelloReuqest) (*vrv_service.Service1HelloResponse, error) {
	response := &vrv_service.Service1HelloResponse{}
	return response, nil
}

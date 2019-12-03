package flexapiapi

import (
	"context"
	flexapi "flexapi/gen/flexapi"
	"log"
)

// flexapi service example implementation.
// The example methods log the requests and return zero values.
type flexapisrvc struct {
	logger *log.Logger
}

// NewFlexapi returns the flexapi service implementation.
func NewFlexapi(logger *log.Logger) flexapi.Service {
	return &flexapisrvc{logger}
}

// Add implements add.
func (s *flexapisrvc) Add(ctx context.Context, p *flexapi.AddPayload) (res int, err error) {
	s.logger.Print("flexapi.add")
	return p.A+p.B,nil
}

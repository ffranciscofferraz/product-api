package server

import (
	"context"

	protos "github.com/franciscofferraz/coffee-shop/currency/protos/currency"
	gohclog "github.com/hashicorp/go-hclog"
)

type Currency struct {
	log gohclog.Logger
}

func NewCurrency(l gohclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

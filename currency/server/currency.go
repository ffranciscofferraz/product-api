package server

import (
	"context"

	"github.com/franciscofferraz/coffee-shop/currency/data"
	protos "github.com/franciscofferraz/coffee-shop/currency/protos/currency"
	gohclog "github.com/hashicorp/go-hclog"
)

type Currency struct {
	rates *data.ExchangeRates
	log   gohclog.Logger
}

func NewCurrency(r *data.ExchangeRates, l gohclog.Logger) *Currency {
	return &Currency{r, l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

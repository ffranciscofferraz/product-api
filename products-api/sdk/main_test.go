package main

import (
	"fmt"
	"testing"

	"github.com/franciscofferraz/coffee-shop/products-api/sdk/client"
	"github.com/franciscofferraz/coffee-shop/products-api/sdk/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	client := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()

	product, err := client.Products.ListProducts(params)

	if err != nil {
		t.Error(err)
	}

	fmt.Println(product)

}

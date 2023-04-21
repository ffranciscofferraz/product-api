package handlers

import "github.com/franciscofferraz/coffee-shop/data"

type errorResponseWrapper struct {
	Body GenericError
}

type errorValidationWrapper struct {
	Body ValidationError
}

type productsResponseWrapper struct {
	Body []data.Product
}

type productResponseWrapper struct {
	Body data.Product
}

type noContentResponseWrapper struct {
}

type productParamsWrapper struct {
	Body data.Product
}

type productIDParamsWrapper struct {
	ID int `json:"id"`
}

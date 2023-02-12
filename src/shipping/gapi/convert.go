package gapi

import (
	"github.com/Arthur199212/microservices-demo/src/shipping/pb"
	"github.com/Arthur199212/microservices-demo/src/shipping/shipping"
)

func convertToProducts(products []*pb.Product) []shipping.Product {
	sp := make([]shipping.Product, len(products))
	for i, p := range products {
		sp[i] = shipping.Product{
			ID:       p.GetId(),
			Quantity: p.GetQuantity(),
		}
	}
	return sp
}

func convertToAddress(address *pb.Address) shipping.Address {
	var state *string = &address.State
	if *state == "" {
		state = nil
	}
	return shipping.Address{
		StreetAddress: address.GetStreetAddress(),
		City:          address.GetCity(),
		Country:       address.GetCountry(),
		ZipCode:       address.GetZipCode(),
		State:         state,
	}
}

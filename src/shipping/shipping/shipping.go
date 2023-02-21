package shipping

import (
	"math"

	"github.com/google/uuid"
)

type ShippingService interface {
	GetQuote(address Address, products []Product) (Quote, error)
	ShipOrder(address Address, products []Product) (string, error)
}

type shippingService struct{}

func NewShippingService() ShippingService {
	return &shippingService{}
}

type Quote struct {
	Quote        float32
	CurrencyCode string
}

const eurCurrencyCode = "EUR"

var EmptyQuote = Quote{
	Quote:        0,
	CurrencyCode: eurCurrencyCode,
}

type Address struct {
	StreetAddress string  `validate:"required,min=5,max=64"`
	City          string  `validate:"required,min=2,max=64"`
	Country       string  `validate:"required,min=2,max=64"`
	ZipCode       string  `validate:"required,numeric,min=4,max=10"`
	State         *string `validate:"omitempty,min=2,max=64"`
}

type Product struct {
	ID       int64 `validate:"required,min=1"`
	Quantity int32 `validate:"required,min=1"`
}

const mockShipmentCostForOneItem = 1.24

func (s *shippingService) GetQuote(address Address, products []Product) (Quote, error) {
	// mock for demo purposes
	var quote float32 = 0
	for _, product := range products {
		quote += float32(product.Quantity) * mockShipmentCostForOneItem
	}
	// formats to '%.2f'
	formatedQuote := float32(math.Round(float64(quote)*100) / 100)

	return Quote{
		Quote:        formatedQuote,
		CurrencyCode: eurCurrencyCode,
	}, nil
}

func (s *shippingService) ShipOrder(address Address, products []Product) (string, error) {
	// mock for demo purposes
	trackingId := uuid.New().String()
	return trackingId, nil
}

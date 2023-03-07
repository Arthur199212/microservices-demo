package models

type Address struct {
	StreetAddress string  `json:"streetAddress" validate:"required,min=5,max=64"`
	City          string  `json:"city" validate:"required,min=2,max=64"`
	Country       string  `json:"country" validate:"required,min=2,max=64"`
	ZipCode       string  `json:"zipCode" validate:"required,numeric,min=4,max=10"`
	State         *string `json:"state,omitempty" validate:"omitempty,min=2,max=64"`
}

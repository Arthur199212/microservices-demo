package db

type Cart struct {
	SessionID string
	Products  []*Product
}

type Product struct {
	ID       int64
	Quantity int32
}

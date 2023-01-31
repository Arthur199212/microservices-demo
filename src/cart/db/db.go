package db

import (
	"database/sql"
	"sync"
)

type CartDB interface {
	AddItem(sessionID string, productID int64, quantity int32) error
	GetCart(sessionID string) (*Cart, error)
	ClearCart(sessionID string) error
}

type FakeStore struct {
	mu    sync.RWMutex
	carts []*Cart
}

func NewCartDB() CartDB {
	return &FakeStore{}
}

func (s *FakeStore) AddItem(
	sessionID string,
	productID int64,
	quantity int32,
) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cart, err := s.findCart(sessionID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// create a new cart
	if err == sql.ErrNoRows {
		cart = &Cart{
			SessionID: sessionID,
			Products: []*Product{
				{
					ID:       productID,
					Quantity: quantity,
				},
			},
		}
		s.carts = append(s.carts, cart)
		return nil
	}

	// edit existing cart
	for _, product := range cart.Products {
		if product.ID == productID {
			// update existing record with product
			product.Quantity = quantity
			return nil
		}
	}
	// add product & quantity if it wasn't there yet
	cart.Products = append(cart.Products, &Product{
		ID: productID, Quantity: quantity,
	})

	return nil
}

func (s *FakeStore) GetCart(sessionID string) (*Cart, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.findCart(sessionID)
}

func (s *FakeStore) ClearCart(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, cart := range s.carts {
		if cart.SessionID == sessionID {
			s.carts[i] = s.carts[len(s.carts)-1]
			s.carts = s.carts[:len(s.carts)-1]
			return nil
		}
	}
	return sql.ErrNoRows
}

// linear search just for demo purposes
func (s *FakeStore) findCart(sessionID string) (*Cart, error) {
	for _, cart := range s.carts {
		if cart.SessionID == sessionID {
			return cart, nil
		}
	}
	return &Cart{}, sql.ErrNoRows
}

package db

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartDB(t *testing.T) {
	sessionId := "mock-session-id"
	db := NewCartDB()
	_, err := db.GetCart(sessionId)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)

	var productId int64 = 1
	var quantity int32 = 5
	err = db.AddItem(sessionId, productId, quantity)
	assert.NoError(t, err)

	cart, err := db.GetCart(sessionId)
	assert.NoError(t, err)
	assert.Equal(t, sessionId, cart.SessionID)
	assert.Len(t, cart.Products, 1)
	assert.Equal(t, productId, cart.Products[0].ID)
	assert.Equal(t, quantity, cart.Products[0].Quantity)

	var productId2 int64 = 2
	var quantity2 int32 = 3
	err = db.AddItem(sessionId, productId2, quantity2)
	assert.NoError(t, err)

	cart, err = db.GetCart(sessionId)
	assert.NoError(t, err)
	assert.Len(t, cart.Products, 2)

	err = db.AddItem(sessionId, productId, quantity+10)
	assert.NoError(t, err)
	cart, err = db.GetCart(sessionId)
	assert.NoError(t, err)
	var product *Product
	for _, p := range cart.Products {
		if p.ID != productId {
			continue
		}
		product = p
		break
	}
	assert.NotNil(t, product)
	assert.Equal(t, quantity+10, product.Quantity)

	err = db.ClearCart(sessionId)
	assert.NoError(t, err)

	_, err = db.GetCart(sessionId)
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}

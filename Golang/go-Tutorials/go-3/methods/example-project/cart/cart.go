package cart

import (
	"os/user"
	"time"

	"first/product"

	"github.com/Rhymond/go-money"
)

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Items        []Item
	CurrencyCode string
	isLocked     bool
}

type Item struct {
	product.Product
	Quantity uint8
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	return nil, nil
}

func (c *Cart) Lock() error {
	return nil
}

func (c *Cart) delete() error {
	return nil
}

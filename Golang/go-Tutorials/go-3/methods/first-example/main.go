package main

import (
	"log"
	"os/user"
	"time"

	"github.com/Rhymond/go-money"
)

type Item struct {
	ID string
}

type Cart struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	lockedAt  time.Time
	user.User
	Item         []Item
	CurrencyCode string
	isLocked     bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
	return nil, nil
}

func (c *Cart) Lock() error {
	return nil
}

func main() {
	newCart := Cart{}

	totalPrice, err := newCart.TotalPrice()
	if err != nil {
		log.Printf("impossible to compute price of the cart: %s", err)
		return
	}
	log.Println("Total Price", totalPrice.Display())

	err = newCart.Lock()
	if err != nil {
		log.Printf("impossible to lock the cart: %s", err)
		return
	}
}

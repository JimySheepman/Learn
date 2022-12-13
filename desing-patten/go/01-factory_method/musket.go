package main

// Concrete Creator
//
// Concrete Creators override the base factory method so it returns a different type of product.
// Note that the factory method doesn’t have to create new instances all the time.
// It can also return existing objects from a cache, an object pool, or another source.
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 3,
		},
	}
}

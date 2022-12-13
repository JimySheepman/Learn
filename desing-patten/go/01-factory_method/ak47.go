package main

// Concrete Creator
//
// Concrete Creators override the base factory method so it returns a different type of product.
// Note that the factory method doesn’t have to create new instances all the time.
// It can also return existing objects from a cache, an object pool, or another source.
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

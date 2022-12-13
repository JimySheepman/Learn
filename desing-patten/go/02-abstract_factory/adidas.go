package main

// Concrete Factory
//
// Concrete Factories implement creation methods of the abstract factory.
// Each concrete factory corresponds to a specific variant of products
// and creates only those product variants.
type Adidas struct {
}

func (a *Adidas) makeShoe() IShirt {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

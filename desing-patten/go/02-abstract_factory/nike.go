package main

// Concrete Factory
//
// Concrete Factories implement creation methods of the abstract factory.
// Each concrete factory corresponds to a specific variant of products
// and creates only those product variants.
type Nike struct {
}

func (n *Nike) makeShoe() IShirt {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 14,
		},
	}
}

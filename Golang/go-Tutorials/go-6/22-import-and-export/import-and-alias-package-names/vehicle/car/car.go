package car

type Car struct {
	Price int
}

func (obj *Car) Single(price int) int {
	obj.Price = price
	return obj.Price
}

func (obj *Car) Double(price int) int {
	obj.Price = price * 2
	return obj.Price
}

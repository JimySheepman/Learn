package price

func totalPrice(nights, rate, cityTax uint) uint {
	return nights*rate + cityTax
}

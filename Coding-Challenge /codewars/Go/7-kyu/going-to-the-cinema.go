package kata

import "math"

func Movie(card, ticket int, perc float64) (count int) {
	var systemA float64
	systemB := float64(card)
	for systemA <= math.Ceil(systemB) {
		count++
		systemA = float64(ticket * count)
		systemB += float64(ticket) * math.Pow(perc, float64(count))
	}
	return count
}

package occupancy

const highLimit = 70.0
const mediumLimit = 20.0

func level(occupancyRate float32) string {
	if occupancyRate > highLimit {
		return "High"
	} else if occupancyRate > mediumLimit {
		return "Medium"
	} else {
		return "Low"
	}
}
func rate(roomOccupied int, totalRooms int) float32 {
	return (float32(roomOccupied) / float32(totalRooms)) * 100
}

package occupancy

const highLimit = 70.0
const mediumLimit = 20.0

func level(occupancyRate float32) string {
	if occupancyRate > highLimit {
		return "high"
	} else if occupancyRate > mediumLimit {
		return "medium"
	} else {
		return "low"
	}
}

func rate(roomsOccupied int, totalRooms int) float32 {
	return (float32(roomsOccupied) / float32(totalRooms)) * 100
}

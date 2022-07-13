package kata

func FindUniq(arr []float32) float32 {
	keys := make(map[float32]bool)
	list := []float32{}
	var st float32
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		} else {
			st = entry
		}
	}
	for i, v := range list {
		if v == st {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list[0]
}

/* Best Practice
func FindUniq(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for i, v := range arr[1:] {
		if v != arr[i] {
			return v
		}
	}
	return 0
}
*/

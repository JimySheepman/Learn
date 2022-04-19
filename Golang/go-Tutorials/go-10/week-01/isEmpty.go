package week_01

func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" || obj == false || obj == 0 {
		return true
	}
	return false
}

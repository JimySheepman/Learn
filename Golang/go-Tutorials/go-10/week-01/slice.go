package week_01

import "reflect"

func Last(arr interface{}) interface{} {
	value := reflect.ValueOf(arr)

	return value.Index(value.Len() - 1).Interface()
}

package slice

import (
	"reflect"
)

// Reverse reverses the order of a slice.
// Keep in mind that this function uses reflection,
// if performance is your priority, please make your own function.
func Reverse(slice interface{}) {
	value := reflect.ValueOf(slice)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Slice {
		return
	}

	temp := reflect.New(value.Index(0).Type()).Elem()
	for i, j := 0, value.Len()-1; i < j; i, j = i+1, j-1 {
		temp.Set(value.Index(i))
		value.Index(i).Set(value.Index(j))
		value.Index(j).Set(temp)
	}
}

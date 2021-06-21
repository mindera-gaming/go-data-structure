package slice

import (
	"reflect"
)

// Reverse reverses the order of a slice.
func Reverse(slice interface{}) {
	value := reflect.ValueOf(slice)
	if value.IsNil() || value.Kind() != reflect.Slice {
		return
	}

	temp := reflect.New(value.Index(0).Type()).Elem()
	for i, j := 0, value.Len()-1; i < j; i, j = i+1, j-1 {
		temp.Set(value.Index(i))
		value.Index(i).Set(value.Index(j))
		value.Index(j).Set(temp)
	}
}

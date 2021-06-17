package array

// Reverse reverses the order of an array.
func Reverse(array *[]interface{}) {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

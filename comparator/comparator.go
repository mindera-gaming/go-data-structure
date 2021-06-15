package comparator

// Compare compares two values.
// Used to sort elements by their value.
type Compare func(value1 interface{}, value2 interface{}) Result

// Result represents the result of the comparison.
type Result byte

const (
	Less    Result = iota // the first value is smaller than the second
	Equal                 // the values are equal
	Greater               // the first value is greater than the second
)

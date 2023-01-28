package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		v, _ := strconv.Atoi(fnb.Value())
		return v
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	prefix := "This is a fancy box containing the number"
	value := float64(ExtractFancyNumber(fnb))
	return fmt.Sprintf("%s %.1f", prefix, value)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var result string
	switch v := i.(type) {
	case int:
		result = DescribeNumber(float64(v))
	case float64:
		result = DescribeNumber(v)
	case NumberBox:
		result = DescribeNumberBox(v)
	case FancyNumberBox:
		result = DescribeFancyNumberBox(v)
	default:
		result = "Return to sender"
	}
	return result
}

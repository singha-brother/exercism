package grains

import "errors"

func Square(n int) (uint64, error) {
	if n <= 0 || n >= 65 {
		return 0, errors.New("number must be (0-64)")
	}
	return 1 << (n - 1), nil
}

func Total() uint64 {
	// it doesn't use above square function
	return (1 << 64) - 1
}

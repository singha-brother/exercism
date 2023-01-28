package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

// Define the Clock type here.
func New(h, m int) Clock {
	for m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h %= 24
		h += 24
	}
	min := m % 60
	hr := h + (m / 60)
	hr = hr % 24
	return Clock{
		hour:   hr,
		minute: min,
	}
}

func (c Clock) Add(m int) Clock {
	c.minute += m
	return New(c.hour, c.minute)
}

func (c Clock) Subtract(m int) Clock {
	c.minute -= m
	return New(c.hour, c.minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

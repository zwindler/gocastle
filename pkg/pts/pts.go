package pts

import "fmt"

type Point int

// New returns a new Point with a given value.
func New(value int) *Point {
	p := Point(value)
	return &p
}

// Add adds a given value to the current point.
func (p *Point) Add(value int) {
	*p += Point(value)
}

// Sub subtracts a given value to the current point.
func (p *Point) Sub(value int) {
	*p -= Point(value)
}

// Reset resets the current point to 0.
func (p *Point) Reset() {
	*p = 0
}

// Get returns the current point value.
func (p *Point) Get() int {
	return int(*p)
}

// Set sets the current point value.
func (p *Point) Set(value int) {
	*p = Point(value)
}

// IsPositive returns true if the current point is greater than 0.
func (p *Point) IsPositive() bool {
	return *p > 0
}

// String returns the current point as a string.
func (p *Point) String() string {
	return fmt.Sprintf("%d", p.Get())
}

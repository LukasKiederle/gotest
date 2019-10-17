package rational

import "fmt"

// Rational struct
type Rational struct {
	numerator   int
	denominator int
}

// Multiply 2 Rationals
func (x Rational) Multiply(y Rational) Rational {
	return NewRational(x.numerator*y.numerator, x.denominator*y.denominator)
}

// NewRational constructor
func NewRational(numerator, denominator int) Rational {
	if denominator == 0 {
		panic("Can't divide through zero!\n")
	}

	r := Rational{}
	divisor := gcd(numerator, denominator)
	r.numerator = numerator / divisor
	r.denominator = denominator / divisor
	return r
}

// Add 2 Rationals
func (x Rational) Add(y Rational) Rational {
	return NewRational(x.numerator*y.denominator+y.numerator*x.denominator, x.denominator*y.denominator)
}

//greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (x Rational) String() string {
	return fmt.Sprintf("(%v/%v)", x.numerator, x.denominator)
}

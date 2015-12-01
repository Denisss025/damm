// Package damm supports the computation of a decimal checksum. It uses a
// method proposed by H. Michael Damm in 2004. The checksum doesn't change for
// leading zeroes.
//
// The function CheckDigit computes the check sum as string:
//
//   c, err := CheckDigit("12345678901")
//
// The function Validate checks whether the appended check digit is correct.
//
//   ok := Validate("123456789018")
//
// Information about the algorithm is available on Wikipedia.
//
// http://en.wikipedia.org/wiki/Damm_algorithm
//
package damm

import (
	"errors"
)

var (
	ErrNonDigitSymbol = errors.New("Digit strings must contain digits only")

	// quasi contains the quasi group used for computing the check digit.
	quasi = [100]int8{
		00, 30, 10, 70, 50, 90, 80, 60, 40, 20,
		70, 00, 90, 20, 10, 50, 40, 80, 60, 30,
		40, 20, 00, 60, 80, 70, 10, 30, 50, 90,
		10, 70, 50, 00, 90, 80, 30, 40, 20, 60,
		60, 10, 20, 30, 00, 40, 50, 90, 70, 80,
		30, 60, 70, 40, 20, 00, 90, 50, 80, 10,
		50, 80, 60, 90, 70, 20, 00, 10, 30, 40,
		80, 90, 40, 50, 30, 60, 20, 00, 10, 70,
		90, 40, 30, 80, 60, 10, 70, 20, 00, 50,
		20, 50, 80, 10, 40, 30, 60, 70, 90, 00,
	}

	toStr = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

// checkInt computes a check digit and returns it as an integer. If the digit
// argument contains non-digit symbols, an error is returned.
func checkInt(digits string) (int, error) {
	var c int8
	for _, x := range digits {
		if x < '0' || x > '9' {
			return 0, ErrNonDigitSymbol
		}
		c = quasi[c+int8(x&0xF)]
	}
	return int(c) / 10, nil
}

// CheckDigit computes the check digit and returns it as a string. The function
// argument must contain decimal digits only.
func CheckDigit(digits string) (string, error) {
	x, err := checkInt(digits)
	return toStr[x], err
}

// Validate checks a number with the check digit appended. The function returns
// true only if the argument contains only decimal digits and the appended
// check digit is correct.
func Validate(digits string) bool {
	x, err := checkInt(digits)
	return err == nil && x == 0
}

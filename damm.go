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
	quasi = [10][10]int8{
		{0, 3, 1, 7, 5, 9, 8, 6, 4, 2},
		{7, 0, 9, 2, 1, 5, 4, 8, 6, 3},
		{4, 2, 0, 6, 8, 7, 1, 3, 5, 9},
		{1, 7, 5, 0, 9, 8, 3, 4, 2, 6},
		{6, 1, 2, 3, 0, 4, 5, 9, 7, 8},
		{3, 6, 7, 4, 2, 0, 9, 5, 8, 1},
		{5, 8, 6, 9, 7, 2, 0, 1, 3, 4},
		{8, 9, 4, 5, 3, 6, 2, 0, 1, 7},
		{9, 4, 3, 8, 6, 1, 7, 2, 0, 5},
		{2, 5, 8, 1, 4, 3, 6, 7, 9, 0},
	}

	toStr = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

// checkInt computes a check digit and returns it as an integer. If the digit
// argument contains non-digit symbols, an error is returned.
func checkInt(digits string) (int, error) {
	var c int8
	for _, x := range digits {
		if !('0' <= x && x <= '9') {
			return 0, ErrNonDigitSymbol
		}
		c = quasi[c][x&0xF]
	}
	return int(c), nil
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

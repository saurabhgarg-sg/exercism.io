// Package leap implements functions for leap years
package leap

// IsLeapYear return true if given year is a leap year, false otherwise
func IsLeapYear(year int) bool {
	if (year % 4 == 0.0 && year % 100 != 0.0) || (year % 400 == 0.0 ) {
		return true
	}
	return false
}

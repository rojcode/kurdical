package kurdical

import "fmt"

// ErrorInvalidYear represents an error for invalid year in Jalaali calculations.
type ErrorInvalidYear struct {
	Year int
}

func (e *ErrorInvalidYear) Error() string {
	return fmt.Sprintf("invalid year: %d", e.Year)
}

// ErrorInvalidMonth represents an error for invalid month.
type ErrorInvalidMonth struct {
	Month int
}

func (e *ErrorInvalidMonth) Error() string {
	return fmt.Sprintf("invalid month: %d", e.Month)
}

// ErrorInvalidDay represents an error for invalid day.
type ErrorInvalidDay struct {
	Day int
}

func (e *ErrorInvalidDay) Error() string {
	return fmt.Sprintf("invalid day: %d", e.Day)
}

// ErrorInvalidDate represents an error for invalid date combination.
type ErrorInvalidDate struct {
	Year  int
	Month int
	Day   int
}

func (e *ErrorInvalidDate) Error() string {
	return fmt.Sprintf("invalid date: year=%d, month=%d, day=%d", e.Year, e.Month, e.Day)
}

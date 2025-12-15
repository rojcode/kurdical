package kurdical

import (
	"fmt"
	"time"
)

// Dialect represents the Kurdish dialect for month names.
type Dialect int

// Epoch represents the historical origin for Kurdish year calculation.
type Epoch int

const (
	// Laki dialect
	Laki Dialect = iota
	// Hawrami dialect
	Hawrami
	// Sorani dialect
	Sorani
	// Kalhuri dialect
	Kalhuri
	// Kurmanji dialect
	Kurmanji
)

const (
	// MedianKingdom epoch (Diako)
	MedianKingdom Epoch = iota
	// FallOfNineveh epoch (Cyaxares)
	FallOfNineveh
)

// KurdishDate represents a date in the Kurdish calendar.
type KurdishDate struct {
	Year      int
	Month     int
	Day       int
	MonthName string
	Dialect   Dialect
	Epoch     Epoch
}

// epochOffsets maps epochs to their year offsets from Solar Hijri.
var epochOffsets = map[Epoch]int{
	MedianKingdom: 1321,
	FallOfNineveh: 1233,
}

// GregorianToKurdish converts a Gregorian time.Time to a KurdishDate.
func GregorianToKurdish(t time.Time, dialect Dialect, epoch Epoch) KurdishDate {
	year, month, day := t.Date()
	sYear, sMonth, sDay := gregorianToSolarHijri(year, int(month), day)
	kYear := sYear + epochOffsets[epoch]
	monthName := monthNames[dialect][sMonth-1]
	return KurdishDate{
		Year:      kYear,
		Month:     sMonth,
		Day:       sDay,
		MonthName: monthName,
		Dialect:   dialect,
		Epoch:     epoch,
	}
}

// KurdishToGregorian converts a KurdishDate to a Gregorian time.Time.
func KurdishToGregorian(k KurdishDate) (time.Time, error) {
	if k.Month < 1 || k.Month > 12 {
		return time.Time{}, fmt.Errorf("invalid month: %d", k.Month)
	}
	monthDays := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
	sYear := k.Year - epochOffsets[k.Epoch]
	if isSolarHijriLeap(sYear) {
		monthDays[11] = 30
	}
	if k.Day < 1 || k.Day > monthDays[k.Month-1] {
		return time.Time{}, fmt.Errorf("invalid day: %d for month %d in year %d", k.Day, k.Month, k.Year)
	}
	gYear, gMonth, gDay := solarHijriToGregorian(sYear, k.Month, k.Day)
	return time.Date(gYear, time.Month(gMonth), gDay, 0, 0, 0, 0, time.UTC), nil
}

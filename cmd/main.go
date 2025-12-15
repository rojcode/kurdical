package main

import (
	"fmt"
	"time"

	"github.com/rojcode/kurdical"
)

func main() {
	// Convert 24 Dec 2025 to Kurdish
	date := time.Date(2025, 10, 11, 0, 0, 0, 0, time.UTC)
	kurdish := kurdical.GregorianToKurdish(date, kurdical.Sorani, kurdical.MedianKingdom)

	fmt.Printf("Gregorian: %s\n", date.Format("2006-01-02"))
	fmt.Printf("Kurdish (Median Kingdom, Sorani): %d-%d-%d %s\n", kurdish.Year, kurdish.Month, kurdish.Day, kurdish.MonthName)

	// Also show Fall of Nineveh epoch
	kurdish2 := kurdical.GregorianToKurdish(date, kurdical.Kurmanji, kurdical.FallOfNineveh)
	fmt.Printf("Kurdish (Fall of Nineveh, Kurmanji): %d-%d-%d %s\n", kurdish2.Year, kurdish2.Month, kurdish2.Day, kurdish2.MonthName)
}

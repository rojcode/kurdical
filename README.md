# kurdical

[![Go](https://github.com/rojcode/kurdical/actions/workflows/go.yml/badge.svg)](https://github.com/rojcode/kurdical/actions/workflows/go.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/rojcode/kurdical)](https://goreportcard.com/report/github.com/rojcode/kurdical)

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue)](https://golang.org/dl/)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## About the Kurdish Calendar

The Kurdish calendar is an independent solar calendar system with historical offsets aligned to significant events in Kurdish history.

It supports two main historical epochs:

- **Median Kingdom (Diako) epoch**: Starting from the establishment of the Median kingdom by Diako
- **Fall of Nineveh (Cyaxares) epoch**: Starting from the fall of Nineveh by Cyaxares

The Kurdish year is calculated by adding epoch-specific offsets to the base solar calendar year.

## Features

- Convert Gregorian dates to Kurdish calendar
- Convert Kurdish calendar dates back to Gregorian
- Support for two historical epochs (Median Kingdom and Fall of Nineveh)
- Month names in 5 Kurdish dialects: Laki, Hawrami, Sorani, Kalhuri, Kurmanji
- Display weekdays in Kurdish
- Format dates with Kurdish digits (٠ ١ ٢ ٣ ٤ ٥ ٦ ٧ ٨ ٩)
- Error handling and date validation
- 100% test coverage
- Complete documentation with practical examples

## Installation

```bash
go get github.com/rojcode/kurdical
```

## Usage

### 1. Basic Gregorian to Kurdish Conversion

```go
package main

import (
    "fmt"
    "time"
    "github.com/rojcode/kurdical"
)

func main() {
    t := time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC)
    k := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.MedianKingdom)
    fmt.Printf("Kurdish date: %d-%d-%d %s\n", k.Year, k.Month, k.Day, k.MonthName)
    // Output: Kurdish date: 2723-1-1 خاکه‌لێوه
}
```

### 2. Round Trip Conversion

```go
g, err := kurdical.KurdishToGregorian(k)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("Gregorian date: %s\n", g.Format("2006-01-02"))
    // Output: Gregorian date: 2023-03-21
}
```

### 3. Different Kurdish Dialects

```go
// Sorani dialect
kSorani := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("Sorani: %s\n", kSorani.MonthName) // خاکه‌لێوه

// Kurmanji dialect
kKurmanji := kurdical.GregorianToKurdish(t, kurdical.Kurmanji, kurdical.MedianKingdom)
fmt.Printf("Kurmanji: %s\n", kKurmanji.MonthName) // نیسان

// Laki dialect
kLaki := kurdical.GregorianToKurdish(t, kurdical.Laki, kurdical.MedianKingdom)
fmt.Printf("Laki: %s\n", kLaki.MonthName) // په‌نجه

// Hawrami dialect
kHawrami := kurdical.GregorianToKurdish(t, kurdical.Hawrami, kurdical.MedianKingdom)
fmt.Printf("Hawrami: %s\n", kHawrami.MonthName) // نه‌ورۆز

// Kalhuri dialect
kKalhuri := kurdical.GregorianToKurdish(t, kurdical.Kalhuri, kurdical.MedianKingdom)
fmt.Printf("Kalhuri: %s\n", kKalhuri.MonthName) // جه‌ژنان (جه‌شنان)
```

### 4. Different Historical Epochs

```go
// Median Kingdom epoch
kMedian := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("Median Kingdom: %d\n", kMedian.Year) // 2723

// Fall of Nineveh epoch
kNineveh := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.FallOfNineveh)
fmt.Printf("Fall of Nineveh: %d\n", kNineveh.Year) // 2635
```

### 5. Weekday Display

```go
fmt.Printf("Weekday: %s\n", kurdical.WeekdayNames[k.Weekday])
// Output: Weekday: سێ‌شەممە
```

### 6. Formatting with Kurdish Digits

```go
formatted, err := k.KFormat("2006-01-02 January")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Formatted: %s\n", formatted)
    // Output: Formatted: ٢٧٢٣-٠١-٠١ خاکه‌لێوه
}
```

### 7. Using Date-Only Functions

```go
k := kurdical.GregorianToKurdishDate(2023, 3, 21, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("Kurdish: %d-%d-%d %s\n", k.Year, k.Month, k.Day, k.MonthName)
```

### 8. Converting Kurdish to Gregorian

```go
gy, gm, gd, err := kurdical.KurdishToGregorianDate(k.Year, k.Month, k.Day, k.Epoch)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Gregorian: %d-%d-%d\n", gy, gm, gd)
}
```

### 9. Creating Kurdish Date Manually

```go
specificKurdish := kurdical.KurdishDate{
    Year:    2725,
    Month:   9,
    Day:     24,
    Dialect: kurdical.Sorani,
    Epoch:   kurdical.MedianKingdom,
}
gregorianSpecific, err := kurdical.KurdishToGregorian(specificKurdish)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Kurdish 2725-09-24 to Gregorian: %s\n", gregorianSpecific.Format("2006-01-02"))
}
```

### 10. Invalid Month Error Handling

```go
// Invalid month
invalidKurdish := kurdical.KurdishDate{
    Year:  2725,
    Month: 13, // Invalid month
    Day:   24,
    Epoch: kurdical.MedianKingdom,
}
_, err := kurdical.KurdishToGregorian(invalidKurdish)
if err != nil {
    fmt.Printf("Error: %s\n", err) // Error: invalid month: 13
}
```

### 11. Invalid Day Error Handling

```go
// Invalid day
invalidDay := kurdical.KurdishDate{
    Year:  2725,
    Month: 1,
    Day:   32, // Invalid day
    Epoch: kurdical.MedianKingdom,
}
_, err := kurdical.KurdishToGregorian(invalidDay)
if err != nil {
    fmt.Printf("Error: %s\n", err) // Error: invalid day: 32 for month 1 in year 2725
}
```

### 12. Different Dates of Year

```go
// January date
janDate := kurdical.GregorianToKurdishDate(2023, 1, 1, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("January: %d-%d-%d %s\n", janDate.Year, janDate.Month, janDate.Day, janDate.MonthName)

// December date
decDate := kurdical.GregorianToKurdishDate(2023, 12, 31, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("December: %d-%d-%d %s\n", decDate.Year, decDate.Month, decDate.Day, decDate.MonthName)
```

### 13. Comparing Epochs

```go
date := time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)

median := kurdical.GregorianToKurdish(date, kurdical.Sorani, kurdical.MedianKingdom)
nineveh := kurdical.GregorianToKurdish(date, kurdical.Sorani, kurdical.FallOfNineveh)

fmt.Printf("Median Kingdom: %d-%d-%d\n", median.Year, median.Month, median.Day)
fmt.Printf("Fall of Nineveh: %d-%d-%d\n", nineveh.Year, nineveh.Month, nineveh.Day)
fmt.Printf("Year difference: %d\n", median.Year - nineveh.Year) // 88
```

### 14. Different Format Layouts

```go
layouts := []string{
    "2006-01-02",
    "2006/01/02 January",
    "02 Jan 2006",
    "Monday, 02 January 2006",
}

for _, layout := range layouts {
    formatted, _ := k.KFormat(layout)
    fmt.Printf("Format %s: %s\n", layout, formatted)
}
```

### 15. Complete Program Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/rojcode/kurdical"
)

func main() {
    // Current date
    now := time.Now()
    kurdishNow := kurdical.GregorianToKurdish(now, kurdical.Sorani, kurdical.MedianKingdom)

    fmt.Printf("Current Gregorian: %s\n", now.Format("2006-01-02 15:04:05"))
    fmt.Printf("Current Kurdish: %d-%d-%d %s (%s)\n",
        kurdishNow.Year, kurdishNow.Month, kurdishNow.Day,
        kurdishNow.MonthName, kurdical.WeekdayNames[kurdishNow.Weekday])

    // Convert to specific Kurdish date
    newYearKurdish := kurdical.KurdishDate{
        Year:  2726,
        Month: 1,
        Day:   1,
        Epoch: kurdical.MedianKingdom,
    }
    newYearGregorian, _ := kurdical.KurdishToGregorian(newYearKurdish)
    fmt.Printf("Kurdish New Year 2726: %s\n", newYearGregorian.Format("2006-01-02"))
}
```

## API

### Types

- `Dialect`: Enum for Kurdish dialects (Laki, Hawrami, Sorani, Kalhuri, Kurmanji)
- `Epoch`: Enum for historical epochs (MedianKingdom, FallOfNineveh)
- `KurdishDate`: Struct representing a date in the Kurdish calendar

### Functions

- `GregorianToKurdish(t time.Time, dialect Dialect, epoch Epoch) KurdishDate`
- `GregorianToKurdishDate(year, month, day int, dialect Dialect, epoch Epoch) KurdishDate`
- `KurdishToGregorian(k KurdishDate) (time.Time, error)`
- `KurdishToGregorianDate(kYear, kMonth, kDay int, epoch Epoch) (int, int, int, error)`
- `(k KurdishDate) KFormat(layout string) (string, error)`: Formats the Kurdish date using Go time layout strings with Kurdish digits

## Kurdish Calendar Details

The Kurdish calendar uses epoch-specific historical adjustments:

- Median Kingdom epoch: Kurdish year = Base year + 1321
- Fall of Nineveh epoch: Kurdish year = Base year + 1233

## Cultural Notes

This module respects Kurdish cultural heritage by providing accurate month names in authentic dialects. UTF-8 encoding ensures proper display of Kurdish characters.

## License

See LICENSE file.

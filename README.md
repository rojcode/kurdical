# kurdical

A Go module for Kurdish calendar utilities.

This module provides conversion between Gregorian and Kurdish calendars, supporting two historical epochs and month names in five Kurdish dialects.

## Features

- Convert Gregorian dates to Kurdish calendar
- Convert Kurdish calendar dates back to Gregorian
- Support for two Kurdish historical epochs:
  - Median Kingdom (Diako)
  - Fall of Nineveh (Cyaxares)
- Month names in 5 Kurdish dialects: Laki, Hawrami, Sorani, Kalhuri, Kurmanji

## Installation

```bash
go get github.com/rojcode/kurdical
```

## Usage

### Basic Conversion

```go
package main

import (
    "fmt"
    "time"
    "github.com/rojcode/kurdical"
)

func main() {
    // Convert Gregorian to Kurdish
    t := time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC)
    k := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.MedianKingdom)
    fmt.Printf("Kurdish date: %d-%d-%d %s\n", k.Year, k.Month, k.Day, k.MonthName)
    // Output: Kurdish date: 2723-1-1 خاکه‌لێوه
}
```

### Round Trip Conversion

```go
// Convert Kurdish to Gregorian
g, err := kurdical.KurdishToGregorian(k)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("Gregorian date: %s\n", g.Format("2006-01-02"))
    // Output: Gregorian date: 2023-03-21
}
```

### Different Dialects

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
```

### Different Epochs

```go
// Median Kingdom epoch
kMedian := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("Median Kingdom: %d\n", kMedian.Year) // 2723

// Fall of Nineveh epoch
kNineveh := kurdical.GregorianToKurdish(t, kurdical.Sorani, kurdical.FallOfNineveh)
fmt.Printf("Fall of Nineveh: %d\n", kNineveh.Year) // 2635
```

### Formatting with Kurdish Digits

```go
formatted, err := k.KFormat("2006-01-02 January")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Formatted: %s\n", formatted)
    // Output: Formatted: ٢٧٢٣-٠١-٠١ خاکه‌لێوه
}
```

### Weekday Display

```go
fmt.Printf("Weekday: %s\n", kurdical.WeekdayNames[k.Weekday])
// Output: Weekday: سێشەممە (for Tuesday)
```

### Creating Kurdish Date Manually

```go
// Create a specific Kurdish date and convert to Gregorian
specificKurdish := kurdical.KurdishDate{
    Year:    2725,
    Month:   9,
    Day:     24,
    Dialect: kurdical.Sorani, // Optional
    Epoch:   kurdical.MedianKingdom,
}
gregorianSpecific, err := kurdical.KurdishToGregorian(specificKurdish)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Kurdish 2725-09-24 to Gregorian: %s\n", gregorianSpecific.Format("2006-01-02"))
    // Output: Kurdish 2725-09-24 to Gregorian: 2025-12-15
}
```

### Date-Only Functions

```go
// Using date-only functions without time.Time
k := kurdical.GregorianToKurdishDate(2023, 3, 21, kurdical.Sorani, kurdical.MedianKingdom)
fmt.Printf("Kurdish: %d-%d-%d %s\n", k.Year, k.Month, k.Day, k.MonthName)

gy, gm, gd, err := kurdical.KurdishToGregorianDate(k.Year, k.Month, k.Day, k.Epoch)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Printf("Gregorian: %d-%d-%d\n", gy, gm, gd)
    // Output: Gregorian: 2023-3-21
}
```

### Error Handling

```go
// Invalid date
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

The Kurdish calendar is based on the Solar Hijri calendar with adjusted epochs.

- Median Kingdom epoch: Kurdish year = Solar Hijri year + 1321
- Fall of Nineveh epoch: Kurdish year = Solar Hijri year + 1233

## Cultural Notes

This module respects Kurdish cultural heritage by providing accurate month names in authentic dialects. The UTF-8 encoding ensures proper display of Kurdish characters.

## License

See LICENSE file.

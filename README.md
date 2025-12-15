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

    // Convert Kurdish to Gregorian
    g, err := kurdical.KurdishToGregorian(k)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Gregorian date: %s\n", g.Format("2006-01-02"))
    }
}
```

## API

### Types

- `Dialect`: Enum for Kurdish dialects (Laki, Hawrami, Sorani, Kalhuri, Kurmanji)
- `Epoch`: Enum for historical epochs (MedianKingdom, FallOfNineveh)
- `KurdishDate`: Struct representing a date in the Kurdish calendar

### Functions

- `GregorianToKurdish(t time.Time, dialect Dialect, epoch Epoch) KurdishDate`
- `KurdishToGregorian(k KurdishDate) (time.Time, error)`

## Kurdish Calendar Details

The Kurdish calendar is based on the Solar Hijri calendar with adjusted epochs.

- Median Kingdom epoch: Kurdish year = Solar Hijri year + 1321
- Fall of Nineveh epoch: Kurdish year = Solar Hijri year + 1233

## Cultural Notes

This module respects Kurdish cultural heritage by providing accurate month names in authentic dialects. The UTF-8 encoding ensures proper display of Kurdish characters.

## License

See LICENSE file.

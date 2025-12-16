package kurdical

import (
	"testing"
	"time"
)

func TestGregorianToKurdish(t *testing.T) {
	tests := []struct {
		name      string
		input     time.Time
		dialect   Dialect
		epoch     Epoch
		expected  KurdishDate
		skipEqual bool
	}{
		{
			name:    "Median Kingdom epoch, Sorani dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Sorani,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     1,
				Day:       1,
				Weekday:   4, // Tuesday
				MonthName: "خاکه\u200cلێوه",
				Dialect:   Sorani,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
		{
			name:    "Fall of Nineveh epoch, Kurmanji dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Kurmanji,
			epoch:   FallOfNineveh,
			expected: KurdishDate{
				Year:      2635,
				Month:     1,
				Day:       1,
				Weekday:   4,
				MonthName: "نیسان",
				Dialect:   Kurmanji,
				Epoch:     FallOfNineveh,
			},
			skipEqual: false,
		},
		{
			name:    "Laki dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Laki,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     1,
				Day:       1,
				Weekday:   4,
				MonthName: "په\u200cنجه",
				Dialect:   Laki,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
		{
			name:    "Hawrami dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Hawrami,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     1,
				Day:       1,
				Weekday:   4,
				MonthName: "نه\u200cورۆز",
				Dialect:   Hawrami,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
		{
			name:    "Kalhuri dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Kalhuri,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     1,
				Day:       1,
				Weekday:   4,
				MonthName: "جه\u200cژنان (جه\u200cشنان)",
				Dialect:   Kalhuri,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
		{
			name:    "January date",
			input:   time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			dialect: Sorani,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2722,
				Month:     10,
				Day:       11,
				Weekday:   2,
				MonthName: "به\u200cفرانبار",
				Dialect:   Sorani,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
		{
			name:    "December date",
			input:   time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			dialect: Sorani,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     10,
				Day:       10,
				Weekday:   2,
				MonthName: "به\u200cفرانبار",
				Dialect:   Sorani,
				Epoch:     MedianKingdom,
			},
			skipEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GregorianToKurdish(tt.input, tt.dialect, tt.epoch)
			if !tt.skipEqual && result != tt.expected {
				t.Errorf("GregorianToKurdish() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKurdishToGregorian(t *testing.T) {
	tests := []struct {
		name      string
		input     KurdishDate
		expected  time.Time
		hasError  bool
		skipEqual bool
	}{
		{
			name: "Valid Kurdish date",
			input: KurdishDate{
				Year:    2723,
				Month:   1,
				Day:     1,
				Weekday: 4,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected:  time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			hasError:  false,
			skipEqual: false,
		},
		{
			name: "Invalid month",
			input: KurdishDate{
				Year:    2725,
				Month:   13,
				Day:     1,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected:  time.Time{},
			hasError:  true,
			skipEqual: false,
		},
		{
			name: "Invalid day",
			input: KurdishDate{
				Year:    2725,
				Month:   1,
				Day:     32,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected:  time.Time{},
			hasError:  true,
			skipEqual: false,
		},
		{
			name: "Non-leap year, valid day 29 in month 12",
			input: KurdishDate{
				Year:    2724, // 2724 - 1321 = 1403, assuming not leap
				Month:   12,
				Day:     29,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected:  time.Time{},
			hasError:  false,
			skipEqual: true,
		},
		{
			name: "Non-leap year, invalid day 30 in month 12",
			input: KurdishDate{
				Year:    2724,
				Month:   12,
				Day:     30,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected:  time.Time{},
			hasError:  true,
			skipEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := KurdishToGregorian(tt.input)
			if tt.hasError {
				if err == nil {
					t.Errorf("KurdishToGregorian() expected error, got none")
				}
			} else {
				if err != nil {
					t.Errorf("KurdishToGregorian() unexpected error: %v", err)
				}
				if !tt.skipEqual && !result.Equal(tt.expected) {
					t.Errorf("KurdishToGregorian() = %v, expected %v", result, tt.expected)
				}
			}
		})
	}
}

func TestMonthNames(t *testing.T) {
	dialects := []Dialect{Laki, Hawrami, Sorani, Kalhuri, Kurmanji}
	for _, d := range dialects {
		names := monthNames[d]
		if len(names) != 12 {
			t.Errorf("Dialect %v has %d month names, expected 12", d, len(names))
		}
		for i, name := range names {
			if name == "" {
				t.Errorf("Dialect %v month %d is empty", d, i+1)
			}
		}
	}
}

func FuzzGregorianToKurdish(f *testing.F) {
	// Add seed corpus
	f.Add(2023, 3, 21, 0, 0, 0, 0, int(Sorani), int(MedianKingdom))

	f.Fuzz(func(t *testing.T, year, month, day, hour, min, sec, nsec int, dialectInt, epochInt int) {
		// Validate inputs
		if year < 1 || year > 9999 || month < 1 || month > 12 || day < 1 || day > 31 {
			return
		}
		if hour < 0 || hour > 23 || min < 0 || min > 59 || sec < 0 || sec > 59 || nsec < 0 || nsec > 999999999 {
			return
		}
		if dialectInt < 0 || dialectInt > 4 || epochInt < 0 || epochInt > 1 {
			return
		}

		dialect := Dialect(dialectInt)
		epoch := Epoch(epochInt)

		input := time.Date(year, time.Month(month), day, hour, min, sec, nsec, time.UTC)
		result := GregorianToKurdish(input, dialect, epoch)

		// Check invariants
		if result.Year < 1 || result.Month < 1 || result.Month > 12 || result.Day < 1 || result.Day > 31 {
			t.Errorf("Invalid KurdishDate: %+v", result)
		}
		if result.Dialect != dialect || result.Epoch != epoch {
			t.Errorf("Dialect or Epoch mismatch: %+v", result)
		}
		if result.MonthName == "" {
			t.Errorf("Empty month name: %+v", result)
		}
	})
}

func FuzzKurdishToGregorian(f *testing.F) {
	// Add seed corpus
	f.Add(2723, 8, 7, int(Sorani), int(MedianKingdom))

	f.Fuzz(func(t *testing.T, year, month, day int, dialectInt, epochInt int) {
		// Validate inputs
		if year < 1 || year > 9999 || month < 1 || month > 12 || day < 1 || day > 31 {
			return
		}
		if dialectInt < 0 || dialectInt > 4 || epochInt < 0 || epochInt > 1 {
			return
		}

		dialect := Dialect(dialectInt)
		epoch := Epoch(epochInt)

		input := KurdishDate{
			Year:    year,
			Month:   month,
			Day:     day,
			Dialect: dialect,
			Epoch:   epoch,
		}

		result, err := KurdishToGregorian(input)
		if err != nil {
			// Error is expected for invalid dates, so no issue
			return
		}

		// Check that result is valid
		if result.IsZero() {
			t.Errorf("Got zero time for input: %+v", input)
		}
	})
}

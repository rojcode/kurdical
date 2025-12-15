package kurdical

import (
	"testing"
	"time"
)

func TestGregorianToKurdish(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		dialect  Dialect
		epoch    Epoch
		expected KurdishDate
	}{
		{
			name:    "Median Kingdom epoch, Sorani dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Sorani,
			epoch:   MedianKingdom,
			expected: KurdishDate{
				Year:      2723,
				Month:     8,
				Day:       7,
				MonthName: "خه‌زه‌ڵوه‌ر",
				Dialect:   Sorani,
				Epoch:     MedianKingdom,
			},
		},
		{
			name:    "Fall of Nineveh epoch, Kurmanji dialect",
			input:   time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC),
			dialect: Kurmanji,
			epoch:   FallOfNineveh,
			expected: KurdishDate{
				Year:      2635,
				Month:     8,
				Day:       7,
				MonthName: "مژدار",
				Dialect:   Kurmanji,
				Epoch:     FallOfNineveh,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GregorianToKurdish(tt.input, tt.dialect, tt.epoch)
			if result != tt.expected {
				t.Errorf("GregorianToKurdish() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestKurdishToGregorian(t *testing.T) {
	tests := []struct {
		name     string
		input    KurdishDate
		expected time.Time
		hasError bool
	}{
		{
			name: "Valid Kurdish date",
			input: KurdishDate{
				Year:    2723,
				Month:   8,
				Day:     7,
				Dialect: Sorani,
				Epoch:   MedianKingdom,
			},
			expected: time.Date(2023, 3, 20, 0, 0, 0, 0, time.UTC),
			hasError: false,
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
			expected: time.Time{},
			hasError: true,
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
			expected: time.Time{},
			hasError: true,
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
				if !result.Equal(tt.expected) {
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

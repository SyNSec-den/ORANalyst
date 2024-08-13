// Implementation of TOML's local date/time.
//
// Copied over from Google's civil to avoid pulling all the Google dependencies.
// Originals:
//   https://raw.githubusercontent.com/googleapis/google-cloud-go/ed46f5086358513cf8c25f8e3f022cb838a49d66/civil/civil.go
// Changes:
//   * Renamed files from civil* to localtime*.
//   * Package changed from civil to toml.
//   * 'Local' prefix added to all structs.
//
// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// Package civil implements types for civil time, a time-zone-independent
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// representation of time that follows the rules of the proleptic
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// Gregorian calendar with exactly 24-hour days, 60-minute hours, and 60-second
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// minutes.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// Because they lack location information, these types do not represent unique
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:25
// moments or intervals of time. Use time.Time for that purpose.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
package toml

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
import (
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:32
)

import (
	"fmt"
	"time"
)

// A LocalDate represents a date (year, month, day).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:39
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:39
// This type does not include location information, and therefore does not
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:39
// describe a unique 24-hour timespan.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:43
type LocalDate struct {
	Year	int		// Year (e.g., 2014).
	Month	time.Month	// Month of the year (January = 1, ...).
	Day	int		// Day of the month, starting at 1.
}

// LocalDateOf returns the LocalDate in which a time occurs in that time's location.
func LocalDateOf(t time.Time) LocalDate {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:50
	_go_fuzz_dep_.CoverTab[123043]++
											var d LocalDate
											d.Year, d.Month, d.Day = t.Date()
											return d
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:53
	// _ = "end of CoverTab[123043]"
}

// ParseLocalDate parses a string in RFC3339 full-date format and returns the date value it represents.
func ParseLocalDate(s string) (LocalDate, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:57
	_go_fuzz_dep_.CoverTab[123044]++
											t, err := time.Parse("2006-01-02", s)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:59
		_go_fuzz_dep_.CoverTab[123046]++
												return LocalDate{}, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:60
		// _ = "end of CoverTab[123046]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:61
		_go_fuzz_dep_.CoverTab[123047]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:61
		// _ = "end of CoverTab[123047]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:61
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:61
	// _ = "end of CoverTab[123044]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:61
	_go_fuzz_dep_.CoverTab[123045]++
											return LocalDateOf(t), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:62
	// _ = "end of CoverTab[123045]"
}

// String returns the date in RFC3339 full-date format.
func (d LocalDate) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:66
	_go_fuzz_dep_.CoverTab[123048]++
											return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:67
	// _ = "end of CoverTab[123048]"
}

// IsValid reports whether the date is valid.
func (d LocalDate) IsValid() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:71
	_go_fuzz_dep_.CoverTab[123049]++
											return LocalDateOf(d.In(time.UTC)) == d
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:72
	// _ = "end of CoverTab[123049]"
}

// In returns the time corresponding to time 00:00:00 of the date in the location.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
// In is always consistent with time.LocalDate, even when time.LocalDate returns a time
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
// on a different day. For example, if loc is America/Indiana/Vincennes, then both
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//	time.LocalDate(1955, time.May, 1, 0, 0, 0, 0, loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
// and
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//	civil.LocalDate{Year: 1955, Month: time.May, Day: 1}.In(loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
// return 23:00:00 on April 30, 1955.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:75
// In panics if loc is nil.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:85
func (d LocalDate) In(loc *time.Location) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:85
	_go_fuzz_dep_.CoverTab[123050]++
											return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:86
	// _ = "end of CoverTab[123050]"
}

// AddDays returns the date that is n days in the future.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:89
// n can also be negative to go into the past.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:91
func (d LocalDate) AddDays(n int) LocalDate {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:91
	_go_fuzz_dep_.CoverTab[123051]++
											return LocalDateOf(d.In(time.UTC).AddDate(0, 0, n))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:92
	// _ = "end of CoverTab[123051]"
}

// DaysSince returns the signed number of days between the date and s, not including the end day.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:95
// This is the inverse operation to AddDays.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:97
func (d LocalDate) DaysSince(s LocalDate) (days int) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:97
	_go_fuzz_dep_.CoverTab[123052]++

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:100
	deltaUnix := d.In(time.UTC).Unix() - s.In(time.UTC).Unix()
											return int(deltaUnix / 86400)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:101
	// _ = "end of CoverTab[123052]"
}

// Before reports whether d1 occurs before d2.
func (d1 LocalDate) Before(d2 LocalDate) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:105
	_go_fuzz_dep_.CoverTab[123053]++
											if d1.Year != d2.Year {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:106
		_go_fuzz_dep_.CoverTab[123056]++
												return d1.Year < d2.Year
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:107
		// _ = "end of CoverTab[123056]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:108
		_go_fuzz_dep_.CoverTab[123057]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:108
		// _ = "end of CoverTab[123057]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:108
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:108
	// _ = "end of CoverTab[123053]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:108
	_go_fuzz_dep_.CoverTab[123054]++
											if d1.Month != d2.Month {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:109
		_go_fuzz_dep_.CoverTab[123058]++
												return d1.Month < d2.Month
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:110
		// _ = "end of CoverTab[123058]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:111
		_go_fuzz_dep_.CoverTab[123059]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:111
		// _ = "end of CoverTab[123059]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:111
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:111
	// _ = "end of CoverTab[123054]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:111
	_go_fuzz_dep_.CoverTab[123055]++
											return d1.Day < d2.Day
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:112
	// _ = "end of CoverTab[123055]"
}

// After reports whether d1 occurs after d2.
func (d1 LocalDate) After(d2 LocalDate) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:116
	_go_fuzz_dep_.CoverTab[123060]++
											return d2.Before(d1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:117
	// _ = "end of CoverTab[123060]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:120
// The output is the result of d.String().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:122
func (d LocalDate) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:122
	_go_fuzz_dep_.CoverTab[123061]++
											return []byte(d.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:123
	// _ = "end of CoverTab[123061]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:126
// The date is expected to be a string in a format accepted by ParseLocalDate.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:128
func (d *LocalDate) UnmarshalText(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:128
	_go_fuzz_dep_.CoverTab[123062]++
											var err error
											*d, err = ParseLocalDate(string(data))
											return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:131
	// _ = "end of CoverTab[123062]"
}

// A LocalTime represents a time with nanosecond precision.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
// This type does not include location information, and therefore does not
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
// describe a unique moment in time.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
// This type exists to represent the TIME type in storage-based APIs like BigQuery.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:134
// Most operations on Times are unlikely to be meaningful. Prefer the LocalDateTime type.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:141
type LocalTime struct {
	Hour		int	// The hour of the day in 24-hour format; range [0-23]
	Minute		int	// The minute of the hour; range [0-59]
	Second		int	// The second of the minute; range [0-59]
	Nanosecond	int	// The nanosecond of the second; range [0-999999999]
}

// LocalTimeOf returns the LocalTime representing the time of day in which a time occurs
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:148
// in that time's location. It ignores the date.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:150
func LocalTimeOf(t time.Time) LocalTime {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:150
	_go_fuzz_dep_.CoverTab[123063]++
											var tm LocalTime
											tm.Hour, tm.Minute, tm.Second = t.Clock()
											tm.Nanosecond = t.Nanosecond()
											return tm
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:154
	// _ = "end of CoverTab[123063]"
}

// ParseLocalTime parses a string and returns the time value it represents.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:157
// ParseLocalTime accepts an extended form of the RFC3339 partial-time format. After
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:157
// the HH:MM:SS part of the string, an optional fractional part may appear,
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:157
// consisting of a decimal point followed by one to nine decimal digits.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:157
// (RFC3339 admits only one digit after the decimal point).
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:162
func ParseLocalTime(s string) (LocalTime, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:162
	_go_fuzz_dep_.CoverTab[123064]++
											t, err := time.Parse("15:04:05.999999999", s)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:164
		_go_fuzz_dep_.CoverTab[123066]++
												return LocalTime{}, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:165
		// _ = "end of CoverTab[123066]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:166
		_go_fuzz_dep_.CoverTab[123067]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:166
		// _ = "end of CoverTab[123067]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:166
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:166
	// _ = "end of CoverTab[123064]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:166
	_go_fuzz_dep_.CoverTab[123065]++
											return LocalTimeOf(t), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:167
	// _ = "end of CoverTab[123065]"
}

// String returns the date in the format described in ParseLocalTime. If Nanoseconds
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:170
// is zero, no fractional part will be generated. Otherwise, the result will
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:170
// end with a fractional part consisting of a decimal point and nine digits.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:173
func (t LocalTime) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:173
	_go_fuzz_dep_.CoverTab[123068]++
											s := fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
											if t.Nanosecond == 0 {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:175
		_go_fuzz_dep_.CoverTab[123070]++
												return s
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:176
		// _ = "end of CoverTab[123070]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:177
		_go_fuzz_dep_.CoverTab[123071]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:177
		// _ = "end of CoverTab[123071]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:177
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:177
	// _ = "end of CoverTab[123068]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:177
	_go_fuzz_dep_.CoverTab[123069]++
											return s + fmt.Sprintf(".%09d", t.Nanosecond)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:178
	// _ = "end of CoverTab[123069]"
}

// IsValid reports whether the time is valid.
func (t LocalTime) IsValid() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:182
	_go_fuzz_dep_.CoverTab[123072]++

											tm := time.Date(2, 2, 2, t.Hour, t.Minute, t.Second, t.Nanosecond, time.UTC)
											return LocalTimeOf(tm) == t
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:185
	// _ = "end of CoverTab[123072]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:188
// The output is the result of t.String().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:190
func (t LocalTime) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:190
	_go_fuzz_dep_.CoverTab[123073]++
											return []byte(t.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:191
	// _ = "end of CoverTab[123073]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:194
// The time is expected to be a string in a format accepted by ParseLocalTime.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:196
func (t *LocalTime) UnmarshalText(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:196
	_go_fuzz_dep_.CoverTab[123074]++
											var err error
											*t, err = ParseLocalTime(string(data))
											return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:199
	// _ = "end of CoverTab[123074]"
}

// A LocalDateTime represents a date and time.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:202
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:202
// This type does not include location information, and therefore does not
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:202
// describe a unique moment in time.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:206
type LocalDateTime struct {
	Date	LocalDate
	Time	LocalTime
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:213
// LocalDateTimeOf returns the LocalDateTime in which a time occurs in that time's location.
func LocalDateTimeOf(t time.Time) LocalDateTime {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:214
	_go_fuzz_dep_.CoverTab[123075]++
											return LocalDateTime{
		Date:	LocalDateOf(t),
		Time:	LocalTimeOf(t),
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:218
	// _ = "end of CoverTab[123075]"
}

// ParseLocalDateTime parses a string and returns the LocalDateTime it represents.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
// ParseLocalDateTime accepts a variant of the RFC3339 date-time format that omits
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
// the time offset but includes an optional fractional time, as described in
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
// ParseLocalTime. Informally, the accepted format is
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
//	YYYY-MM-DDTHH:MM:SS[.FFFFFFFFF]
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:221
// where the 'T' may be a lower-case 't'.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:227
func ParseLocalDateTime(s string) (LocalDateTime, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:227
	_go_fuzz_dep_.CoverTab[123076]++
											t, err := time.Parse("2006-01-02T15:04:05.999999999", s)
											if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:229
		_go_fuzz_dep_.CoverTab[123078]++
												t, err = time.Parse("2006-01-02t15:04:05.999999999", s)
												if err != nil {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:231
			_go_fuzz_dep_.CoverTab[123079]++
													return LocalDateTime{}, err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:232
			// _ = "end of CoverTab[123079]"
		} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:233
			_go_fuzz_dep_.CoverTab[123080]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:233
			// _ = "end of CoverTab[123080]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:233
		}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:233
		// _ = "end of CoverTab[123078]"
	} else {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:234
		_go_fuzz_dep_.CoverTab[123081]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:234
		// _ = "end of CoverTab[123081]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:234
	}
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:234
	// _ = "end of CoverTab[123076]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:234
	_go_fuzz_dep_.CoverTab[123077]++
											return LocalDateTimeOf(t), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:235
	// _ = "end of CoverTab[123077]"
}

// String returns the date in the format described in ParseLocalDate.
func (dt LocalDateTime) String() string {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:239
	_go_fuzz_dep_.CoverTab[123082]++
											return dt.Date.String() + "T" + dt.Time.String()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:240
	// _ = "end of CoverTab[123082]"
}

// IsValid reports whether the datetime is valid.
func (dt LocalDateTime) IsValid() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:244
	_go_fuzz_dep_.CoverTab[123083]++
											return dt.Date.IsValid() && func() bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:245
		_go_fuzz_dep_.CoverTab[123084]++
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:245
		return dt.Time.IsValid()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:245
		// _ = "end of CoverTab[123084]"
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:245
	}()
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:245
	// _ = "end of CoverTab[123083]"
}

// In returns the time corresponding to the LocalDateTime in the given location.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// If the time is missing or ambigous at the location, In returns the same
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// result as time.LocalDate. For example, if loc is America/Indiana/Vincennes, then
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// both
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//	time.LocalDate(1955, time.May, 1, 0, 30, 0, 0, loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// and
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//	civil.LocalDateTime{
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//	    civil.LocalDate{Year: 1955, Month: time.May, Day: 1}},
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//	    civil.LocalTime{Minute: 30}}.In(loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// return 23:30:00 on April 30, 1955.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
//
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:248
// In panics if loc is nil.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:261
func (dt LocalDateTime) In(loc *time.Location) time.Time {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:261
	_go_fuzz_dep_.CoverTab[123085]++
											return time.Date(dt.Date.Year, dt.Date.Month, dt.Date.Day, dt.Time.Hour, dt.Time.Minute, dt.Time.Second, dt.Time.Nanosecond, loc)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:262
	// _ = "end of CoverTab[123085]"
}

// Before reports whether dt1 occurs before dt2.
func (dt1 LocalDateTime) Before(dt2 LocalDateTime) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:266
	_go_fuzz_dep_.CoverTab[123086]++
											return dt1.In(time.UTC).Before(dt2.In(time.UTC))
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:267
	// _ = "end of CoverTab[123086]"
}

// After reports whether dt1 occurs after dt2.
func (dt1 LocalDateTime) After(dt2 LocalDateTime) bool {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:271
	_go_fuzz_dep_.CoverTab[123087]++
											return dt2.Before(dt1)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:272
	// _ = "end of CoverTab[123087]"
}

// MarshalText implements the encoding.TextMarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:275
// The output is the result of dt.String().
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:277
func (dt LocalDateTime) MarshalText() ([]byte, error) {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:277
	_go_fuzz_dep_.CoverTab[123088]++
											return []byte(dt.String()), nil
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:278
	// _ = "end of CoverTab[123088]"
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:281
// The datetime is expected to be a string in a format accepted by ParseLocalDateTime
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:283
func (dt *LocalDateTime) UnmarshalText(data []byte) error {
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:283
	_go_fuzz_dep_.CoverTab[123089]++
											var err error
											*dt, err = ParseLocalDateTime(string(data))
											return err
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:286
	// _ = "end of CoverTab[123089]"
}

//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:287
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/pelletier/go-toml@v1.9.4/localtime.go:287
var _ = _go_fuzz_dep_.CoverTab

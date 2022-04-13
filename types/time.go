/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2022 Albert Moky
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * ==============================================================================
 */
package types

import (
	"math"
	"time"
)

type Time interface {

	// Date returns the year, month, and day in which t occurs.
	Date() (year int, month time.Month, day int)

	// Clock returns the hour, minute, and second within the day specified by t.
	Clock() (hour, min, sec int)

	// Weekday returns the day of the week specified by t.
	Weekday() time.Weekday

	// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
	// and [1,366] in leap years.
	YearDay() int

	// Year returns the year in which t occurs.
	Year() int

	// Month returns the month of the year specified by t.
	Month() time.Month

	// Day returns the day of the month specified by t.
	Day() int

	// Hour returns the hour within the day specified by t, in the range [0, 23].
	Hour() int

	// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
	Minute() int

	// Second returns the second offset within the minute specified by t, in the range [0, 59].
	Second() int

	// Nanosecond returns the nanosecond offset within the second specified by t,
	// in the range [0, 999999999].
	Nanosecond() int

	// Unix returns t as a Unix time, the number of seconds elapsed
	// since January 1, 1970 UTC. The result does not depend on the
	// location associated with t.
	// Unix-like operating systems often record time as a 32-bit
	// count of seconds, but since the method here returns a 64-bit
	// value it is valid for billions of years into the past or future.
	Unix() int64

	// UnixNano returns t as a Unix time, the number of nanoseconds elapsed
	// since January 1, 1970 UTC. The result is undefined if the Unix time
	// in nanoseconds cannot be represented by an int64 (a date before the year
	// 1678 or after 2262). Note that this means the result of calling UnixNano
	// on the zero Time is undefined. The result does not depend on the
	// location associated with t.
	UnixNano() int64

	// IsZero reports whether t represents the zero time instant,
	// January 1, year 1, 00:00:00 UTC.
	IsZero() bool

	// Equal reports whether t and u represent the same time instant.
	// Two times can be equal even if they are in different locations.
	// For example, 6:00 +0200 and 4:00 UTC are Equal.
	// See the documentation on the Time type for the pitfalls of using == with
	// Time values; most code should use Equal instead.
	Equal(u time.Time) bool

	// String returns the time formatted using the format string
	//	"2006-01-02 15:04:05.999999999 -0700 MST"
	//
	// If the time has a monotonic clock reading, the returned string
	// includes a final field "m=±<value>", where value is the monotonic
	// clock reading formatted as a decimal number of seconds.
	//
	// The returned string is meant for debugging; for a stable serialized
	// representation, use t.MarshalText, t.MarshalBinary, or t.Format
	// with an explicit format string.
	String() string

	// Format returns a textual representation of the time value formatted
	// according to layout, which defines the format by showing how the reference
	// time, defined to be
	//	Mon Jan 2 15:04:05 -0700 MST 2006
	// would be displayed if it were the value; it serves as an example of the
	// desired output. The same display rules will then be applied to the time
	// value.
	//
	// A fractional second is represented by adding a period and zeros
	// to the end of the seconds section of layout string, as in "15:04:05.000"
	// to format a time stamp with millisecond precision.
	//
	// Predefined layouts ANSIC, UnixDate, RFC3339 and others describe standard
	// and convenient representations of the reference time. For more information
	// about the formats and the definition of the reference time, see the
	// documentation for ANSIC and the other constants defined by this package.
	Format(layout string) string
}

func TimeIsNil(t Time) bool {
	if ValueIsNil(t) {
		return true
	}
	return t.IsZero()
}

func TimeNil() Time {
	return time.Time{}
}

func TimeNow() Time {
	return time.Now()
}

// timestamp in seconds
func Timestamp(t Time) int64 {
	return t.Unix()
}

// timestamp in nanoseconds
func TimestampNano(t Time) int64 {
	return t.UnixNano()
}

// timestamp in seconds
func TimeToInt64(t Time) int64 {
	return t.Unix()
}

// timestamp in seconds
func TimeToFloat64(t Time) float64 {
	secs := t.Unix()
	nano := t.Nanosecond()
	return float64(secs) + float64(nano) / 1e9
}

func TimeFromInt64(seconds int64) Time {
	return time.Unix(seconds, 0)
}

func TimeFromFloat64(seconds float64) Time {
	trunc, frac := math.Modf(seconds)
	secs := int64(trunc)
	nano := int64(frac * 1e9)
	return time.Unix(secs, nano)
}

// parse from timestamp in seconds
func TimeParse(timestamp interface{}) Time {
	if ValueIsNil(timestamp) {
		return TimeNil()
	}
	return TimeFromFloat64(timestamp.(float64))
}

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

	// IsZero reports whether t represents the zero time instant,
	// January 1, year 1, 00:00:00 UTC.
	IsZero() bool
}

func TimeIsNil(t Time) bool {
	return t.IsZero()
}

func TimeNil() Time {
	return time.Time{}
}

func TimeNow() Time {
	return time.Now()
}

func UnixTime(t Time) int64 {
	return t.Unix()
}

func TimeToInt64(t Time) int64 {
	return t.Unix()
}

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

func TimeParse(timestamp interface{}) Time {
	if ValueIsNil(timestamp) {
		return TimeNil()
	} else {
		return TimeFromFloat64(timestamp.(float64))
	}
}

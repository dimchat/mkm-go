/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2026 Albert Moky
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
	"fmt"
	"strconv"
)

// Default Converter
type DataConverter struct {
	//Converter
}

//
//  String
//

// Override
func (DataConverter) GetString(value any, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	// string
	case string:
		if v == "" {
			return defaultValue
		}
		return v
	// boolean
	case bool:
		return strconv.FormatBool(v)
	// integer
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	// unsigned integer
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	// float number
	case float32:
		return formatFloat(float64(v), 32)
	case float64:
		return formatFloat(v, 64)
	default:
		// unknown type
		return fmt.Sprintf("%v", value)
		//return defaultValue
	}
}

func formatFloat(f float64, bitSize int) string {
	return strconv.FormatFloat(f, 'f', -1, bitSize)
}

//
//  Date Time
//

// Override
func (DataConverter) GetTime(value any, defaultValue Time) Time {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case Time:
		return v
	}
	seconds := ConvertFloat64(value, 0)
	if seconds > 0 {
		return TimeFromFloat64(seconds)
	}
	//panic("Timestamp error: " + ConvertString(value, ""))
	return defaultValue
}

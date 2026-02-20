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
	"errors"
	"strconv"
	"strings"
)

//
//  Float Number
//

// Override
func (DataConverter) GetFloat32(value interface{}, defaultValue float32) float32 {
	v, err := convFloat64(value)
	if err != nil {
		return defaultValue
	}
	return float32(v)
}

// Override
func (DataConverter) GetFloat64(value interface{}, defaultValue float64) float64 {
	v, err := convFloat64(value)
	if err != nil {
		return defaultValue
	}
	return v
}

func convFloat64(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("nil value")
	}
	switch v := value.(type) {
	// float number
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	// integer
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	// unsigned integer
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	// boolean
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	// string
	case string:
		s := strings.TrimSpace(v)
		i, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	default:
		// unknown type
		return 0, errors.New("invalid type")
	}
}

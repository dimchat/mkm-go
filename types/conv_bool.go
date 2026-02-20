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
	"strings"
)

//
//  Boolean
//

//goland:noinspection GoSnakeCaseUsage
var BOOLEAN_STATES = map[string]bool{
	"1": true, "yes": true, "true": true, "on": true,

	"0": false, "no": false, "false": false, "off": false,
	//"+0": false, "-0": false, "0.0": false, "+0.0": false, "-0.0": false,
	"null": false, "none": false, "undefined": false,
}

//goland:noinspection GoSnakeCaseUsage
var MAX_BOOLEAN_LEN = len("undefined")

func parseBool(s string) (bool, error) {
	text := strings.TrimSpace(s)
	size := len(text)
	if size == 0 {
		return false, errors.New("string empty")
	} else if size > MAX_BOOLEAN_LEN {
		return false, errors.New("string too long")
	}
	text = strings.ToLower(text)
	state, exists := BOOLEAN_STATES[text]
	if !exists {
		return false, errors.New("unknown boolean state")
	}
	return state, nil
}

// Override
func (DataConverter) GetBool(value interface{}, defaultValue bool) bool {
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	// boolean
	case bool:
		return v
	// string
	case string:
		b, err := parseBool(v)
		if err != nil {
			return defaultValue
		}
		return b
	// integer
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v != 0
	// float number
	case float32, float64:
		return v != 0.0
	default:
		// unknown type
		return defaultValue
	}
}

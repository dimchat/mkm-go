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
	"reflect"
	"strconv"
	"strings"
)

func parseUint(s string, bitSize int) (uint64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseUint(s, 10, bitSize)
}

// Override
func (DataConverter) GetUInt(value interface{}, defaultValue uint) uint {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return uint(v)
	case int8:
		return uint(v)
	case int16:
		return uint(v)
	case int32:
		return uint(v)
	case int64:
		return uint(v)
	case uint:
		return v
	case uint8:
		return uint(v)
	case uint16:
		return uint(v)
	case uint32:
		return uint(v)
	case uint64:
		return uint(v)
	case float32:
		return uint(v)
	case float64:
		return uint(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := parseUint(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return uint(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		i, err := parseUint(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint(i)
	default:
		//panic(fmt.Sprintf("uint value error: %v", value))
	}
	return defaultValue
}

// Override
func (DataConverter) GetUInt8(value interface{}, defaultValue uint8) uint8 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return uint8(v)
	case int8:
		return uint8(v)
	case int16:
		return uint8(v)
	case int32:
		return uint8(v)
	case int64:
		return uint8(v)
	case uint:
		return uint8(v)
	case uint8:
		return v
	case uint16:
		return uint8(v)
	case uint32:
		return uint8(v)
	case uint64:
		return uint8(v)
	case float32:
		return uint8(v)
	case float64:
		return uint8(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := parseUint(v, 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint8(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint8(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint8(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return uint8(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		i, err := parseUint(rv.String(), 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint8(i)
	default:
		//panic(fmt.Sprintf("uint8 value error: %v", value))
	}
	return defaultValue
}

// Override
func (DataConverter) GetUInt16(value interface{}, defaultValue uint16) uint16 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return uint16(v)
	case int8:
		return uint16(v)
	case int16:
		return uint16(v)
	case int32:
		return uint16(v)
	case int64:
		return uint16(v)
	case uint:
		return uint16(v)
	case uint8:
		return uint16(v)
	case uint16:
		return v
	case uint32:
		return uint16(v)
	case uint64:
		return uint16(v)
	case float32:
		return uint16(v)
	case float64:
		return uint16(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := parseUint(v, 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint16(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint16(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint16(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return uint16(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		i, err := parseUint(rv.String(), 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint16(i)
	default:
		//panic(fmt.Sprintf("uint16 value error: %v", value))
	}
	return defaultValue
}

// Override
func (DataConverter) GetUInt32(value interface{}, defaultValue uint32) uint32 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return uint32(v)
	case int8:
		return uint32(v)
	case int16:
		return uint32(v)
	case int32:
		return uint32(v)
	case int64:
		return uint32(v)
	case uint:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint16:
		return uint32(v)
	case uint32:
		return v
	case uint64:
		return uint32(v)
	case float32:
		return uint32(v)
	case float64:
		return uint32(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := parseUint(v, 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint32(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint32(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint32(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return uint32(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		i, err := parseUint(rv.String(), 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint32(i)
	default:
		//panic(fmt.Sprintf("uint32 value error: %v", value))
	}
	return defaultValue
}

// Override
func (DataConverter) GetUInt64(value interface{}, defaultValue uint64) uint64 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		i, err := parseUint(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint()
	case reflect.Float32, reflect.Float64:
		return uint64(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		i, err := parseUint(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	default:
		//panic(fmt.Sprintf("uint64 value error: %v", value))
	}
	return defaultValue
}

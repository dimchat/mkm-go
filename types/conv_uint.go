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

func (conv *DataConverter) GetUInt(value interface{}, defaultValue uint) uint {
	value = ObjectValue(value)
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case uint:
		return v
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return v.(uint)
	case float32, float64:
		return v.(uint)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseUint(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint(i)
	}
	// other types
	rv := reflect.ValueOf(value)
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
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseUint(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint(i)
	}
	//panic(fmt.Sprintf("uint value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetUInt8(value interface{}, defaultValue uint8) uint8 {
	value = ObjectValue(value)
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case uint8:
		return v
	case int, int8, int16, int32, int64, uint, uint16, uint32, uint64:
		return v.(uint8)
	case float32, float64:
		return v.(uint8)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseUint(v, 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint8(i)
	}
	// other types
	rv := reflect.ValueOf(value)
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
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseUint(rv.String(), 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint8(i)
	}
	//panic(fmt.Sprintf("uint8 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetUInt16(value interface{}, defaultValue uint16) uint16 {
	value = ObjectValue(value)
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case uint16:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint32, uint64:
		return v.(uint16)
	case float32, float64:
		return v.(uint16)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseUint(v, 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint16(i)
	}
	// other types
	rv := reflect.ValueOf(value)
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
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseUint(rv.String(), 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint16(i)
	}
	//panic(fmt.Sprintf("uint16 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetUInt32(value interface{}, defaultValue uint32) uint32 {
	value = ObjectValue(value)
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case uint32:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint64:
		return v.(uint32)
	case float32, float64:
		return v.(uint32)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseUint(v, 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint32(i)
	}
	// other types
	rv := reflect.ValueOf(value)
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
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseUint(rv.String(), 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return uint32(i)
	}
	//panic(fmt.Sprintf("uint32 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetUInt64(value interface{}, defaultValue uint64) uint64 {
	value = ObjectValue(value)
	if value == nil {
		return defaultValue
	}
	switch v := value.(type) {
	case uint64:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32:
		return v.(uint64)
	case float32, float64:
		return v.(uint64)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseUint(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	// other types
	rv := reflect.ValueOf(value)
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
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseUint(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	//panic(fmt.Sprintf("uint64 value error: %v", value))
	return defaultValue
}

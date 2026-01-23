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

func a2i(s string) (int, error) {
	s = strings.TrimSpace(s)
	return strconv.Atoi(s)
}

func parseInt(s string, bitSize int) (int64, error) {
	s = strings.TrimSpace(s)
	return strconv.ParseInt(s, 10, bitSize)
}

func (conv *DataConverter) GetInt(value interface{}, defaultValue int) int {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := a2i(v)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	// other types
	switch rv.Kind() {
	case reflect.Int:
		return int(rv.Int())
	case reflect.Int8:
		return int(rv.Int())
	case reflect.Int16:
		return int(rv.Int())
	case reflect.Int32:
		return int(rv.Int())
	case reflect.Int64:
		return int(rv.Int())
	case reflect.Uint:
		return int(rv.Uint())
	case reflect.Uint8:
		return int(rv.Uint())
	case reflect.Uint16:
		return int(rv.Uint())
	case reflect.Uint32:
		return int(rv.Uint())
	case reflect.Uint64:
		return int(rv.Uint())
	case reflect.Float32:
		return int(rv.Float())
	case reflect.Float64:
		return int(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		i, err := a2i(rv.String())
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	//panic(fmt.Sprintf("int value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetInt8(value interface{}, defaultValue int8) int8 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int8:
		return v
	case int, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v.(int8)
	case float32, float64:
		return v.(int8)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseInt(v, 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int8(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int8(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int8(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return int8(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseInt(rv.String(), 8)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int8(i)
	}
	//panic(fmt.Sprintf("int8 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetInt16(value interface{}, defaultValue int16) int16 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int16:
		return v
	case int, int8, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v.(int16)
	case float32, float64:
		return v.(int16)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseInt(v, 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int16(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int16(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int16(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return int16(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseInt(rv.String(), 16)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int16(i)
	}
	//panic(fmt.Sprintf("int16 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetInt32(value interface{}, defaultValue int32) int32 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int32:
		return v
	case int, int8, int16, int64, uint, uint8, uint16, uint32, uint64:
		return v.(int32)
	case float32, float64:
		return v.(int32)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseInt(v, 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int32(i)
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int32(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int32(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return int32(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseInt(rv.String(), 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return int32(i)
	}
	//panic(fmt.Sprintf("int32 value error: %v", value))
	return defaultValue
}

func (conv *DataConverter) GetInt64(value interface{}, defaultValue int64) int64 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case int64:
		return v
	case int, int8, int16, int32, uint, uint8, uint16, uint32, uint64:
		return v.(int64)
	case float32, float64:
		return v.(int64)
	case bool:
		if v {
			return 1
		} else {
			return 0
		}
	case string:
		i, err := parseInt(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	// other types
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(rv.Uint())
	case reflect.Float32, reflect.Float64:
		return int64(rv.Float())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		} else {
			return 0
		}
	case reflect.String:
		i, err := parseInt(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return i
	}
	//panic(fmt.Sprintf("int64 value error: %v", value))
	return defaultValue
}

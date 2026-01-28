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
	"reflect"
	"strconv"
	"strings"
)

/**
 *  Default Converter
 */
type DataConverter struct{}

//
//  String
//

func i2a(i int) string {
	return strconv.Itoa(i)
}

func formatBool(b bool) string {
	return strconv.FormatBool(b)
}

func formatInt(i int64) string {
	return strconv.FormatInt(i, 10)
}

func formatUint(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func formatFloat(f float64, bitSize int) string {
	return strconv.FormatFloat(f, 'f', -1, bitSize)
}

func formatComplex(c complex128, bitSize int) string {
	return strconv.FormatComplex(c, 'f', -1, bitSize)
}

func (conv *DataConverter) GetString(value interface{}, defaultValue string) string {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case string:
		//if v == "" {
		//	return defaultValue
		//}
		return v
	case bool:
		return formatBool(v)
	case int:
		return i2a(v)
	case int8:
		return formatInt(int64(v))
	case int16:
		return formatInt(int64(v))
	case int32:
		return formatInt(int64(v))
	case int64:
		return formatInt(v)
	case uint:
		return formatUint(uint64(v))
	case uint8:
		return formatUint(uint64(v))
	case uint16:
		return formatUint(uint64(v))
	case uint32:
		return formatUint(uint64(v))
	case uint64:
		return formatUint(v)
	case float32:
		return formatFloat(float64(v), 32)
	case float64:
		return formatFloat(v, 64)
	case complex64:
		//return fmt.Sprintf("%g", v)
		return formatComplex(complex128(v), 64)
	case complex128:
		//return fmt.Sprintf("%g", v)
		return formatComplex(v, 128)
	}
	// other types
	switch rv.Kind() {
	case reflect.String:
		s := rv.String()
		//if s == "" {
		//	return defaultValue
		//}
		return s
	case reflect.Bool:
		return formatBool(rv.Bool())
	case reflect.Int:
		return i2a(int(rv.Int()))
	case reflect.Int8:
		return formatInt(rv.Int())
	case reflect.Int16:
		return formatInt(rv.Int())
	case reflect.Int32:
		return formatInt(rv.Int())
	case reflect.Int64:
		return formatInt(rv.Int())
	case reflect.Uint:
		return formatUint(rv.Uint())
	case reflect.Uint8:
		return formatUint(rv.Uint())
	case reflect.Uint16:
		return formatUint(rv.Uint())
	case reflect.Uint32:
		return formatUint(rv.Uint())
	case reflect.Uint64:
		return formatUint(rv.Uint())
	case reflect.Float32:
		return formatFloat(rv.Float(), 32)
	case reflect.Float64:
		return formatFloat(rv.Float(), 64)
	case reflect.Complex64:
		return formatComplex(rv.Complex(), 64)
	case reflect.Complex128:
		return formatComplex(rv.Complex(), 128)
	default:
		//panic(fmt.Sprintf("string value error: %v", value))
	}
	//return defaultValue
	return fmt.Sprintf("%v", value)
}

func parseBool(s string) (bool, error) {
	text := strings.TrimSpace(s)
	size := len(text)
	if size == 0 {
		return false, &strconv.NumError{Func: "ParseBool", Num: "String empty", Err: strconv.ErrSyntax}
	} else if size > MAX_BOOLEAN_LEN {
		//panic("Bool value error: " + text)
		return false, &strconv.NumError{Func: "ParseBool", Num: "String to long", Err: strconv.ErrRange}
	}
	text = strings.ToLower(text)
	state, exists := BOOLEAN_STATES[text]
	if !exists {
		return false, &strconv.NumError{Func: "ParseBool", Num: "Key not exists", Err: strconv.ErrRange}
	}
	return state, nil
}

func (conv *DataConverter) GetBool(value interface{}, defaultValue bool) bool {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case bool:
		return v
	case string:
		b, err := parseBool(v)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return b
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v != 0
	case float32, float64:
		return v != 0
	case complex64, complex128:
		return v != 0
	}
	// other types
	switch rv.Kind() {
	case reflect.Bool:
		return rv.Bool()
	case reflect.String:
		b, err := parseBool(rv.String())
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return b
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint() != 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return rv.Complex() != 0
	default:
		//panic(fmt.Sprintf("bool value error: %v", value))
	}
	return defaultValue
}

func (conv *DataConverter) GetTime(value interface{}, defaultValue Time) Time {
	target := ObjectTargetValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case Time:
		return v
	}
	ts := ConvertFloat64(value, 0)
	if ts > 0 {
		return TimeFromFloat64(ts)
	}
	//panic("Timestamp error: " + ConvertString(value, ""))
	return defaultValue
}

//
//  Math
//

func parseFloat(s string, bitSize int) (float64, error) {
	text := strings.TrimSpace(s)
	return strconv.ParseFloat(text, bitSize)
}

func parseComplex(s string, bitSize int) (complex128, error) {
	text := strings.TrimSpace(s)
	c, err := strconv.ParseComplex(text, bitSize)
	if err != nil {
		if !strings.ContainsAny(text, "ij") {
			r, err := strconv.ParseFloat(text, 64)
			if err == nil {
				return complex(r, 0), nil
			}
			// "3i"
			if strings.HasSuffix(text, "i") {
				im := text[:len(text)-1]
				if im == "" || im == "+" {
					return complex(0, 1), nil
				} else if im == "-" {
					return complex(0, -1), nil
				}
				b, err := strconv.ParseFloat(im, 64)
				if err == nil {
					return complex(0, b), nil
				}
			}
		}
		return 0, err
	}
	return c, nil
}

func (conv *DataConverter) GetFloat32(value interface{}, defaultValue float32) float32 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case float32:
		return v
	case float64:
		return float32(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v.(float32)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		f, err := parseFloat(v, 32)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return float32(f)
	}
	// other types
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		return float32(rv.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float32(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float32(rv.Uint())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		f, err := parseFloat(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return float32(f)
	default:
		//panic(fmt.Sprintf("float32 value error: %v", value))
	}
	return defaultValue
}

func (conv *DataConverter) GetFloat64(value interface{}, defaultValue float64) float64 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return v.(float64)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		f, err := parseFloat(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return f
	}
	// other types
	switch rv.Kind() {
	case reflect.Float32, reflect.Float64:
		return rv.Float()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(rv.Uint())
	case reflect.Bool:
		if rv.Bool() {
			return 1
		}
		return 0
	case reflect.String:
		f, err := parseFloat(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return f
	default:
		//panic(fmt.Sprintf("float64 value error: %v", value))
	}
	return defaultValue
}

func (conv *DataConverter) GetComplex64(value interface{}, defaultValue complex64) complex64 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case complex64:
		return v
	case complex128:
		return complex64(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return complex64(complex(v.(float64), 0))
	case float32, float64:
		return complex64(complex(v.(float64), 0))
	case string:
		c, err := parseComplex(v, 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return complex64(c)
	}
	// other types
	switch rv.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return complex64(rv.Complex())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex64(complex(float64(rv.Int()), 0))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex64(complex(float64(rv.Uint()), 0))
	case reflect.Float32, reflect.Float64:
		return complex64(complex(rv.Float(), 0))
	case reflect.String:
		c, err := parseComplex(rv.String(), 64)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return complex64(c)
	default:
		//panic(fmt.Sprintf("complex64 value error: %v", value))
	}
	return defaultValue
}

func (conv *DataConverter) GetComplex128(value interface{}, defaultValue complex128) complex128 {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return defaultValue
	}
	switch v := target.(type) {
	case complex64:
		return complex128(v)
	case complex128:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return complex(v.(float64), 0)
	case float32, float64:
		return complex(v.(float64), 0)
	case string:
		c, err := parseComplex(v, 128)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return c
	}
	// other types
	switch rv.Kind() {
	case reflect.Complex64, reflect.Complex128:
		return rv.Complex()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return complex(float64(rv.Int()), 0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return complex(float64(rv.Uint()), 0)
	case reflect.Float32, reflect.Float64:
		return complex(rv.Float(), 0)
	case reflect.String:
		c, err := parseComplex(rv.String(), 128)
		if err != nil {
			//panic(err)
			return defaultValue
		}
		return c
	default:
		//panic(fmt.Sprintf("complex128 value error: %v", value))
	}
	return defaultValue
}

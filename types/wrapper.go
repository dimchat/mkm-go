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
)

/**
 *  Default Data Wrapper
 */
type DataWrapper struct {
	//Wrapper
}

// Override
func (DataWrapper) GetString(value interface{}) string {
	if value == nil {
		//panic(fmt.Sprintf("string value error: %v", value))
		return ""
	}
	switch v := value.(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	}
	// other types
	target, rv := ObjectReflectValue(value)
	if target == nil {
		//panic(fmt.Sprintf("string value error: %v", value))
		return ""
	} else if v, ok := target.(string); ok {
		return v
	}
	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	default:
		//panic(fmt.Sprintf("not a string value: %v", value))
	}
	return fmt.Sprintf("%v", value)
}

// Override
func (DataWrapper) GetMap(value interface{}) StringKeyMap {
	if value == nil {
		//panic(fmt.Sprintf("map value error: %v", value))
		return nil
	}
	switch v := value.(type) {
	case Mapper:
		return v.Map()
	case StringKeyMap:
		return v
	}
	// other types
	target, rv := ObjectReflectValue(value)
	if target == nil {
		//panic(fmt.Sprintf("map value error: %v", value))
		return nil
	} else if v, ok := target.(StringKeyMap); ok {
		return v
	}
	switch rv.Kind() {
	case reflect.Map:
		return reflectMap(rv)
	default:
		//panic(fmt.Sprintf("not a map value: %v", value))
	}
	return nil
}

// Override
func (DataWrapper) GetList(value interface{}) []interface{} {
	if value == nil {
		//panic(fmt.Sprintf("list value error: %v", value))
		return nil
	}
	switch v := value.(type) {
	case []interface{}:
		return v
	}
	// other types
	target, rv := ObjectReflectValue(value)
	if target == nil {
		//panic(fmt.Sprintf("list value error: %v", value))
		return nil
	} else if v, ok := target.([]interface{}); ok {
		return v
	}
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		return reflectList(rv)
	default:
		//panic(fmt.Sprintf("not a list value: %v", value))
	}
	return nil
}

// Override
func (DataWrapper) Unwrap(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case Mapper:
		return UnwrapMap(v.Map())
	case StringKeyMap:
		return UnwrapMap(v)
	case []interface{}:
		return UnwrapList(v)
	case Stringer: // fmt.Stringer:
		return v.String()
	}
	// other types
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return nil
	} else if v, ok := target.(StringKeyMap); ok {
		return UnwrapMap(v)
	} else if v, ok := target.([]interface{}); ok {
		return UnwrapList(v)
	} else if v, ok := target.(string); ok {
		return v
	}
	switch rv.Kind() {
	case reflect.Map:
		return UnwrapMap(reflectMap(rv))
	case reflect.Array, reflect.Slice:
		return UnwrapList(reflectList(rv))
	case reflect.String:
		return rv.String()
	default:
		return value
	}
}

// Override
func (DataWrapper) UnwrapMap(dict StringKeyMap) StringKeyMap {
	if dict == nil {
		return nil
	}
	// unwrap recursively
	result := make(StringKeyMap, len(dict))
	for key, value := range dict {
		result[key] = Unwrap(value)
	}
	return result
}

// Override
func (DataWrapper) UnwrapList(array []interface{}) []interface{} {
	if array == nil {
		return nil
	}
	// unwrap recursively
	result := make([]interface{}, len(array))
	for index, item := range array {
		result[index] = Unwrap(item)
	}
	return result
}

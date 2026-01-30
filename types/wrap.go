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
 *  Data Wrapper
 */
type Wrapper interface {

	/**
	 *  Get inner string
	 *  ~~~~~~~~~~~~~~~~
	 *  Remove first wrapper
	 */
	GetString(str interface{}) string

	/**
	 *  Get inner map
	 *  ~~~~~~~~~~~~~
	 *  Remove first wrapper
	 */
	GetMap(dict interface{}) StringKeyMap

	/**
	 *  Get inner list
	 *  ~~~~~~~~~~~~~~
	 *  Remove first wrapper
	 */
	GetList(array interface{}) []interface{}

	/**
	 *  Unwrap recursively
	 *  ~~~~~~~~~~~~~~~~~~
	 *  Remove all wrappers
	 */
	Unwrap(object interface{}) interface{}

	// Unwrap values for keys in map
	UnwrapMap(dict StringKeyMap) StringKeyMap

	// Unwrap values in the array
	UnwrapList(array []interface{}) []interface{}
}

var sharedWrapper Wrapper = &DataWrapper{}

func SetWrapper(wrapper Wrapper) {
	sharedWrapper = wrapper
}

//
//  Interfaces
//

func FetchString(str interface{}) string {
	return sharedWrapper.GetString(str)
}

func FetchMap(dict interface{}) StringKeyMap {
	return sharedWrapper.GetMap(dict)
}

func FetchList(array interface{}) []interface{} {
	return sharedWrapper.GetList(array)
}

func Unwrap(object interface{}) interface{} {
	return sharedWrapper.Unwrap(object)
}

func UnwrapMap(dict StringKeyMap) StringKeyMap {
	return sharedWrapper.UnwrapMap(dict)
}

func UnwrapList(array []interface{}) []interface{} {
	return sharedWrapper.UnwrapList(array)
}

/**
 *  Default Data Wrapper
 */
type DataWrapper struct {
	//Wrapper
}

// Override
func (wp DataWrapper) GetString(value interface{}) string {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		panic(fmt.Sprintf("string value error: %v", value))
		//return ""
	}
	switch v := target.(type) {
	case fmt.Stringer:
		return v.String()
	case string:
		return v
	}
	// other types
	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	default:
		panic(fmt.Sprintf("not a string value: %v", value))
	}
	//return fmt.Sprintf("%v", value)
}

// Override
func (wp DataWrapper) GetMap(value interface{}) StringKeyMap {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		panic(fmt.Sprintf("map value error: %v", value))
		//return NewMap()
	}
	switch v := target.(type) {
	case Mapper:
		return v.Map()
	case StringKeyMap:
		return v
	}
	// other types
	switch rv.Kind() {
	case reflect.Map:
		return reflectMap(rv)
	default:
		//panic(fmt.Sprintf("not a map value: %v", value))
	}
	return nil
}

// Override
func (wp DataWrapper) GetList(value interface{}) []interface{} {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		panic(fmt.Sprintf("list value error: %v", value))
		//return make([]interface{}, 0)
	}
	switch v := target.(type) {
	case []interface{}:
		return v
	}
	// other types
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		return reflectList(rv)
	default:
		//panic(fmt.Sprintf("not a list value: %v", value))
	}
	return nil
}

// Override
func (wp DataWrapper) Unwrap(value interface{}) interface{} {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return nil
	}
	switch v := target.(type) {
	case Mapper:
		return UnwrapMap(v.Map())
	case StringKeyMap:
		return UnwrapMap(v)
	case []interface{}:
		return UnwrapList(v)
	case fmt.Stringer:
		return v.String()
	}
	// other types
	switch rv.Kind() {
	case reflect.Map:
		return UnwrapMap(reflectMap(rv))
	case reflect.Array, reflect.Slice:
		return UnwrapList(reflectList(rv))
	case reflect.String:
		return rv.String()
	default:
		return target
	}
}

// Override
func (wp DataWrapper) UnwrapMap(value StringKeyMap) StringKeyMap {
	target, rv := ObjectReflectValue(value)
	if target == nil {
		return nil
	}
	// convert to string key map
	var dict StringKeyMap
	switch v := target.(type) {
	case Mapper:
		dict = v.Map()
	case StringKeyMap:
		dict = v
	default:
		// other types
		switch rv.Kind() {
		case reflect.Map:
			dict = reflectMap(rv)
		default:
			//panic(fmt.Sprintf("map value error: %v", value))
			return nil
		}
	}
	// unwrap recursively
	result := NewMap()
	for key, value := range dict {
		result[FetchString(key)] = Unwrap(value)
	}
	return result
}

// Override
func (wp DataWrapper) UnwrapList(array []interface{}) []interface{} {
	result := make([]interface{}, len(array))
	for index, item := range array {
		result[index] = Unwrap(item)
	}
	return result
}

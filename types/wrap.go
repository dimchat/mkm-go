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

import "reflect"

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

/**
 *  Get inner String
 *  <p>
 *      Remove first wrapper
 *  </p>
 */
func FetchString(str interface{}) string {
	return sharedWrapper.GetString(str)
}

/**
 *  Get inner Map
 *  <p>
 *      Remove first wrapper
 *  </p>
 */
func FetchMap(dict interface{}) StringKeyMap {
	return sharedWrapper.GetMap(dict)
}

func FetchList(array interface{}) []interface{} {
	return sharedWrapper.GetList(array)
}

/**
 *  Unwrap recursively
 *  <p>
 *      Remove all wrappers
 *  </p>
 */
func Unwrap(object interface{}) interface{} {
	return sharedWrapper.Unwrap(object)
}

func UnwrapMap(dict StringKeyMap) StringKeyMap {
	return sharedWrapper.UnwrapMap(dict)
}

func UnwrapList(array []interface{}) []interface{} {
	return sharedWrapper.UnwrapList(array)
}

//
//  Reflect Value
//

func reflectMap(rv reflect.Value) StringKeyMap {
	// check type
	dict, ok := rv.Interface().(StringKeyMap)
	if ok {
		return dict
	}
	size := rv.Len()
	dict = make(StringKeyMap, size)
	// copy map entries from reflection
	iter := rv.MapRange()
	for iter.Next() {
		key := iter.Key()
		//if key.Kind() != reflect.String {
		//	//panic(fmt.Sprintf("map key error: %v", key))
		//	continue
		//}
		dict[key.String()] = reflectItemValue(iter.Value())
	}
	return dict
}

func reflectList(rv reflect.Value) []interface{} {
	// check type
	array, ok := rv.Interface().([]interface{})
	if ok {
		return array
	}
	size := rv.Len()
	array = make([]interface{}, size)
	// copy list items from reflection
	for index := 0; index < size; index++ {
		array[index] = reflectItemValue(rv.Index(index))
	}
	return array
}

func reflectItemValue(value reflect.Value) interface{} {
	switch value.Kind() {
	case reflect.Map:
		return reflectMap(value)
	case reflect.Array, reflect.Slice:
		return reflectList(value)
	default:
		return value.Interface()
	}
}

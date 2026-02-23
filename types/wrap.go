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

// Wrapper is a data wrapper interface
//
// Defines various unwrapping methods for wrapped data,
// supporting string, Map, list and other types
type Wrapper interface {

	// GetString gets the string value inside the wrapper
	// Removes the first layer of wrapping and returns the original string
	GetString(str any) string

	// GetMap gets the map value inside the wrapper
	// Removes the first layer of wrapping and returns the original map of StringKeyMap type
	GetMap(dict any) StringKeyMap

	// GetList gets the list value inside the wrapper
	// Removes the first layer of wrapping and returns the original list of []any type
	GetList(array any) []any

	// Unwrap unwraps the object recursively
	// Removes all layers of wrapping and returns the most original object value
	Unwrap(object any) any

	// UnwrapMap unwraps all values corresponding to keys in the Map
	// Performs unwrapping operation on each value in StringKeyMap and returns a new Map
	UnwrapMap(dict StringKeyMap) StringKeyMap

	// UnwrapList unwraps all elements in the list
	// Performs unwrapping operation on each element in []any and returns a new list
	UnwrapList(array []any) []any
}

var sharedWrapper Wrapper = &DataWrapper{}

func SetWrapper(wrapper Wrapper) {
	sharedWrapper = wrapper
}

//
//  Interfaces
//

// FetchString gets the string value inside the wrapper
// Removes the first layer of wrapping and returns the original string
func FetchString(str any) string {
	return sharedWrapper.GetString(str)
}

// FetchMap gets the map value inside the wrapper
// Removes the first layer of wrapping and returns the original map of StringKeyMap type
func FetchMap(dict any) StringKeyMap {
	return sharedWrapper.GetMap(dict)
}

func FetchList(array any) []any {
	return sharedWrapper.GetList(array)
}

// Unwrap unwraps the object recursively
// Removes all layers of wrapping and returns the most original object value
func Unwrap(object any) any {
	return sharedWrapper.Unwrap(object)
}

func UnwrapMap(dict StringKeyMap) StringKeyMap {
	return sharedWrapper.UnwrapMap(dict)
}

func UnwrapList(array []any) []any {
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

func reflectList(rv reflect.Value) []any {
	// check type
	array, ok := rv.Interface().([]any)
	if ok {
		return array
	}
	size := rv.Len()
	array = make([]any, size)
	// copy list items from reflection
	for index := 0; index < size; index++ {
		array[index] = reflectItemValue(rv.Index(index))
	}
	return array
}

func reflectItemValue(value reflect.Value) any {
	switch value.Kind() {
	case reflect.Map:
		return reflectMap(value)
	case reflect.Array, reflect.Slice:
		return reflectList(value)
	default:
		return value.Interface()
	}
}

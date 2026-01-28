/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 Albert Moky
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
)

/**
 *  Mutable Dictionary Wrapper
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *  typedef:
 *      Map<string, *>
 */
type Dictionary struct {
	_dictionary StringKeyMap
}

func (dict *Dictionary) Init(dictionary StringKeyMap) {
	if ValueIsNil(dictionary) {
		// create empty map
		dictionary = NewMap()
	}
	dict._dictionary = dictionary
}

//-------- IObject

func (dict *Dictionary) Equal(other interface{}) bool {
	if other == nil {
		return dict.IsEmpty()
	} else if other == dict {
		// same object
		return true
	}
	// check targeted value
	target, rv := ObjectReflectValue(other)
	if target == nil {
		return dict.IsEmpty()
	}
	// check value types
	switch v := target.(type) {
	case Mapper:
		other = v.Map()
	case StringKeyMap:
		other = v
	default:
		// other types
		switch rv.Kind() {
		case reflect.Map:
			other = reflectMap(rv)
		default:
			// type not matched
			return false
		}
	}
	return reflect.DeepEqual(dict._dictionary, other)
}

//-------- IMap

func (dict *Dictionary) Get(key string) interface{} {
	return dict._dictionary[key]
}

func (dict *Dictionary) Set(key string, value interface{}) {
	if ValueIsNil(value) {
		delete(dict._dictionary, key)
	} else {
		dict._dictionary[key] = value
	}
}

func (dict *Dictionary) Remove(key string) {
	delete(dict._dictionary, key)
}

func (dict *Dictionary) IsEmpty() bool {
	return len(dict._dictionary) == 0
}

func (dict *Dictionary) Keys() []string {
	return MapKeys(dict._dictionary)
}

func (dict *Dictionary) Map() StringKeyMap {
	return dict._dictionary
}

func (dict *Dictionary) CopyMap(deep bool) StringKeyMap {
	if deep {
		return DeepCopyMap(dict._dictionary)
	}
	return CopyMap(dict._dictionary)
}

//-------- Convert values

func (dict *Dictionary) GetString(key string, defaultValue string) string {
	value := dict.Get(key)
	return ConvertString(value, defaultValue)
}

func (dict *Dictionary) GetBool(key string, defaultValue bool) bool {
	value := dict.Get(key)
	return ConvertBool(value, defaultValue)
}

func (dict *Dictionary) GetInt(key string, defaultValue int) int {
	value := dict.Get(key)
	return ConvertInt(value, defaultValue)
}
func (dict *Dictionary) GetInt8(key string, defaultValue int8) int8 {
	value := dict.Get(key)
	return ConvertInt8(value, defaultValue)
}
func (dict *Dictionary) GetInt16(key string, defaultValue int16) int16 {
	value := dict.Get(key)
	return ConvertInt16(value, defaultValue)
}
func (dict *Dictionary) GetInt32(key string, defaultValue int32) int32 {
	value := dict.Get(key)
	return ConvertInt32(value, defaultValue)
}
func (dict *Dictionary) GetInt64(key string, defaultValue int64) int64 {
	value := dict.Get(key)
	return ConvertInt64(value, defaultValue)
}

func (dict *Dictionary) GetUInt(key string, defaultValue uint) uint {
	value := dict.Get(key)
	return ConvertUInt(value, defaultValue)
}
func (dict *Dictionary) GetUInt8(key string, defaultValue uint8) uint8 {
	value := dict.Get(key)
	return ConvertUInt8(value, defaultValue)
}
func (dict *Dictionary) GetUInt16(key string, defaultValue uint16) uint16 {
	value := dict.Get(key)
	return ConvertUInt16(value, defaultValue)
}
func (dict *Dictionary) GetUInt32(key string, defaultValue uint32) uint32 {
	value := dict.Get(key)
	return ConvertUInt32(value, defaultValue)
}
func (dict *Dictionary) GetUInt64(key string, defaultValue uint64) uint64 {
	value := dict.Get(key)
	return ConvertUInt64(value, defaultValue)
}

func (dict *Dictionary) GetFloat32(key string, defaultValue float32) float32 {
	value := dict.Get(key)
	return ConvertFloat32(value, defaultValue)
}
func (dict *Dictionary) GetFloat64(key string, defaultValue float64) float64 {
	value := dict.Get(key)
	return ConvertFloat64(value, defaultValue)
}

func (dict *Dictionary) GetComplex64(key string, defaultValue complex64) complex64 {
	value := dict.Get(key)
	return ConvertComplex64(value, defaultValue)
}
func (dict *Dictionary) GetComplex128(key string, defaultValue complex128) complex128 {
	value := dict.Get(key)
	return ConvertComplex128(value, defaultValue)
}

func (dict *Dictionary) GetTime(key string, defaultValue Time) Time {
	value := dict.Get(key)
	return ConvertTime(value, defaultValue)
}
func (dict *Dictionary) SetTime(key string, value Time) {
	if ValueIsNil(value) {
		dict.Remove(key)
	} else {
		dict.Set(key, TimeToFloat64(value))
	}
}

func (dict *Dictionary) SetStringer(key string, value Stringer) {
	if ValueIsNil(value) {
		dict.Remove(key)
	} else {
		dict.Set(key, value.String())
	}
}

func (dict *Dictionary) SetMapper(key string, value Mapper) {
	if ValueIsNil(value) {
		dict.Remove(key)
	} else {
		dict.Set(key, value.Map())
	}
}

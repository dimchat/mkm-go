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

	_dictionary map[string]interface{}
}

func (dict *Dictionary) Init(dictionary map[string]interface{}) Mapper {
	if ValueIsNil(dictionary) {
		// create empty map
		dictionary = make(map[string]interface{})
	}
	dict._dictionary = dictionary
	return dict
}

//-------- IObject

func (dict *Dictionary) Equal(other interface{}) bool {
	if other == nil {
		return len(dict._dictionary) == 0
	} else if other == dict {
		// same object
		return true
	}
	// check value
	v := ObjectValue(other)
	if v == nil {
		return len(dict._dictionary) == 0
	} else if p, ok := v.(Mapper); ok {
		other = p.Map()
	//} else if p, ok := v.(map[string]interface{}); ok {
	//	other = p
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

func (dict *Dictionary) Keys() []string {
	return MapKeys(dict._dictionary)
}

func (dict *Dictionary) Map() map[string]interface{} {
	return dict._dictionary
}

func (dict *Dictionary) CopyMap(deep bool) map[string]interface{} {
	if deep {
		return DeepCopyMap(dict._dictionary)
	} else {
		return CopyMap(dict._dictionary)
	}
}

//-------- Convert values

func (dict *Dictionary) GetString(key string, defaultValue string) string {
	return ConvertString(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetBool(key string, defaultValue bool) bool {
	return ConvertBool(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetInt(key string, defaultValue int) int {
	return ConvertInt(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetInt8(key string, defaultValue int8) int8 {
	return ConvertInt8(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetInt16(key string, defaultValue int16) int16 {
	return ConvertInt16(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetInt32(key string, defaultValue int32) int32 {
	return ConvertInt32(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetInt64(key string, defaultValue int64) int64 {
	return ConvertInt64(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetUInt(key string, defaultValue uint) uint {
	return ConvertUInt(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetUInt8(key string, defaultValue uint8) uint8 {
	return ConvertUInt8(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetUInt16(key string, defaultValue uint16) uint16 {
	return ConvertUInt16(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetUInt32(key string, defaultValue uint32) uint32 {
	return ConvertUInt32(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetUInt64(key string, defaultValue uint64) uint64 {
	return ConvertUInt64(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetFloat32(key string, defaultValue float32) float32 {
	return ConvertFloat32(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetFloat64(key string, defaultValue float64) float64 {
	return ConvertFloat64(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetComplex64(key string, defaultValue complex64) complex64 {
	return ConvertComplex64(dict._dictionary[key], defaultValue)
}
func (dict *Dictionary) GetComplex128(key string, defaultValue complex128) complex128 {
	return ConvertComplex128(dict._dictionary[key], defaultValue)
}

func (dict *Dictionary) GetTime(key string, defaultValue Time) Time {
	return ConvertTime(dict._dictionary[key], defaultValue)
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

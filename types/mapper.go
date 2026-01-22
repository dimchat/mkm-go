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

type Mapper interface {
	Object

	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)

	/**
	 *  Get all keys
	 */
	Keys() []string

	/**
	 *  Get inner map
	 */
	Map() map[string]interface{}

	/**
	 *  Copy inner map
	 */
	CopyMap(deep bool) map[string]interface{}

	//
	//  Convert values
	//

	GetString    (key string, defaultValue string) string

	GetBool      (key string, defaultValue bool) bool

	GetInt       (key string, defaultValue int) int
	GetInt8      (key string, defaultValue int8) int8
	GetInt16     (key string, defaultValue int16) int16
	GetInt32     (key string, defaultValue int32) int32
	GetInt64     (key string, defaultValue int64) int64

	GetUInt      (key string, defaultValue uint) uint
	GetUInt8     (key string, defaultValue uint8) uint8
	GetUInt16    (key string, defaultValue uint16) uint16
	GetUInt32    (key string, defaultValue uint32) uint32
	GetUInt64    (key string, defaultValue uint64) uint64

	GetFloat32   (key string, defaultValue float32) float32
	GetFloat64   (key string, defaultValue float64) float64

	GetComplex64 (key string, defaultValue complex64) complex64
	GetComplex128(key string, defaultValue complex128) complex128

	GetTime      (key string, defaultValue Time) Time
	SetTime      (key string, value Time)

	SetStringer  (key string, value Stringer)
	SetMapper    (key string, value Mapper)

}

func MapKeys(dictionary map[string]interface{}) []string {
	index := 0
	keys := make([]string, len(dictionary))
	for key := range dictionary {
		keys[index] = key
		index++
	}
	return keys
}

//func FetchMap(dictionary interface{}) map[string]interface{} {
//	wrapper, ok := dictionary.(Mapper)
//	if ok {
//		return wrapper.Map()
//	} else {
//		return dictionary.(map[string]interface{})
//	}
//}

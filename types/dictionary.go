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

type Map interface {
	Object

	Get(key string) interface{}
	Set(key string, value interface{})

	GetMap(clone bool) map[string]interface{}
}

type Dictionary struct {
	Map

	_dictionary map[string]interface{}
}

func (dict *Dictionary) Init(dictionary map[string]interface{}) *Dictionary {
	if dictionary == nil {
		dictionary = make(map[string]interface{})
	}
	dict._dictionary = dictionary
	return dict
}

func (dict *Dictionary) Equal(other interface{}) bool {
	ptr, ok := other.(*Dictionary)
	if !ok {
		obj, ok := other.(Dictionary)
		if !ok {
			return false
		}
		ptr = &obj
	}
	if dict == ptr {
		return true
	}
	// check inner maps
	return MapsEqual(dict._dictionary, ptr._dictionary)
}

func (dict *Dictionary) Get(key string) interface{} {
	return dict._dictionary[key]
}

func (dict *Dictionary) Set(key string, value interface{}) {
	if value == nil {
		delete(dict._dictionary, key)
	} else {
		dict._dictionary[key] = value
	}
}

func (dict *Dictionary) GetMap(clone bool) map[string]interface{} {
	if clone {
		return CloneMap(dict._dictionary)
	} else {
		return dict._dictionary
	}
}

func CloneMap(dictionary map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range dictionary {
		clone[key] = value
	}
	return clone
}

func MapKeys(dictionary map[string]interface{}) []string {
	index := 0
	keys := make([]string, len(dictionary))
	for k := range dictionary {
		keys[index] = k
		index++
	}
	return keys
}

func MapsEqual(map1, map2 map[string]interface{}) bool {
	return reflect.DeepEqual(map1, map2)
}

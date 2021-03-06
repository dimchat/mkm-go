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
	IMap
}
type IMap interface {

	Get(key string) interface{}
	Set(key string, value interface{})
	Remove(key string)

	Keys() []string

	GetMap(clone bool) map[string]interface{}
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

func CloneMap(dictionary map[string]interface{}) map[string]interface{} {
	clone := make(map[string]interface{})
	for key, value := range dictionary {
		clone[key] = value
	}
	return clone
}

/**
 *  Mutable Dictionary Wrapper
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *  typedef:
 *      Map<string, *>
 */
type Dictionary struct {
	BaseObject
	IMap

	_dictionary map[string]interface{}
}

func (dict *Dictionary) Init(dictionary map[string]interface{}) *Dictionary {
	if ValueIsNil(dictionary) {
		dictionary = make(map[string]interface{})
	}
	dict._dictionary = dictionary
	return dict
}

func (dict *Dictionary) Equal(other interface{}) bool {
	// compare pointers
	if dict == other {
		return true
	}
	// compare inner maps
	wrapper, ok := other.(Map)
	if ok {
		return reflect.DeepEqual(dict._dictionary, wrapper.GetMap(false))
	}
	table, ok := other.(map[string]interface{})
	return ok && reflect.DeepEqual(dict._dictionary, table)
}

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

func (dict *Dictionary) GetMap(clone bool) map[string]interface{} {
	if clone {
		return CloneMap(dict._dictionary)
	} else {
		return dict._dictionary
	}
}

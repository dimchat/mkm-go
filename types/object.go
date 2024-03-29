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

type Object interface {

	Equal(other interface{}) bool
}

/**
 *  Base Object
 *  ~~~~~~~~~~~
 */
type BaseObject struct {}

func (obj *BaseObject) Init() Object {
	return obj
}

//-------- IObject

func (obj *BaseObject) Equal(other interface{}) bool {
	value := reflect.ValueOf(other)
	if value.Kind() == reflect.Ptr {
		return obj == other
	} else {
		return *obj == value.Elem().Interface()
	}
}

func ObjectsEqual(i1, i2 interface{}) bool {
	v1 := reflect.ValueOf(i1)
	v2 := reflect.ValueOf(i2)
	if v1.Kind() == reflect.Ptr {
		if v2.Kind() == reflect.Ptr {
			// both i1, i2 are pointers
			return i1 == i2 || v1.Elem().Interface() == v2.Elem().Interface()
		} else {
			// i1 is pointer
			return v1.Elem().Interface() == i2
		}
	} else if v2.Kind() == reflect.Ptr {
		// i2 is pointer
		return i1 == v2.Elem().Interface()
	} else {
		// both i1, i2 are values
		return i1 == i2
	}
}

func ObjectValue(i interface{}) interface{} {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr {
		return value.Elem().Interface()
	} else {
		return i
	}
}

func ObjectIsPointer(i interface{}) bool {
	value := reflect.ValueOf(i)
	return value.Kind() == reflect.Ptr
}

func ObjectPointer(i interface{}) interface{} {
	value := reflect.ValueOf(i)
	if value.Kind() == reflect.Ptr {
		return i
	} else {
		return &i
	}
}

func ValueIsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	defer func() {
		recover()
	}()
	value := reflect.ValueOf(i)
	return value.IsNil()
}

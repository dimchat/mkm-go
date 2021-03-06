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
 *  Base Object
 *  ~~~~~~~~~~~
 */
type Object interface {

	Equal(other interface{}) bool
}

type BaseObject struct {
	Object
}

func (obj *BaseObject) Equal(other interface{}) bool {
	value := reflect.ValueOf(other)
	if value.Kind() == reflect.Ptr {
		return obj == other
	} else {
		return *obj == value.Elem().Interface()
	}
}

/**
 *  Manual Reference Counting
 *  ~~~~~~~~~~~~~~~~~~~~~~~~~
 */
type MRC interface {

	// Set "this" pointer and increase retain count
	Retain(this SelfReference) SelfReference

	// Decrease retain count and return it,
	// if equals 0, erase "this" pointer
	Release() int

	// Append this object to AutoreleasePool
	Autorelease() SelfReference
}

/**
 *  Self Referred Object
 *  ~~~~~~~~~~~~~~~~~~~~
 *
 *  Inheritable
 */
type SelfReference interface {
	MRC

	// Get "this" pointer
	Self() SelfReference

	//Super() Object
}

type InheritableObject struct {
	SelfReference

	_this SelfReference
	_retainCount int
}

func (obj *InheritableObject) Init() *InheritableObject {
	obj._this = nil
	obj._retainCount = 1
	return obj
}

func (obj *InheritableObject) Retain(this SelfReference) SelfReference {
	if this != nil {
		obj._this = this
	}
	obj._retainCount++
	return obj._this
}
func (obj *InheritableObject) Release() int {
	obj._retainCount--
	if obj._retainCount == 0 {
		// break circular reference
		obj._this = nil
	} else if obj._retainCount < 0 {
		panic(obj)
	}
	return obj._retainCount
}
func (obj *InheritableObject) Autorelease() SelfReference {
	return AutoreleasePoolAppend(obj)
}

func (obj *InheritableObject) Self() SelfReference {
	return obj._this
}
//func (obj *BaseObject) Super() Object {
//	panic("super empty")
//}

//--------

/**
 *  Set "this" pointer and increase retain count
 */
func ObjectRetain(obj interface{}) SelfReference {
	o, ok := obj.(SelfReference)
	if ok {
		// call 'Retain()' from child class
		s := o.Self()
		if s == nil {
			return o.Retain(o)
		} else {
			return s.Retain(s)
		}
	} else {
		return o
	}
}

/**
 *  Decrease retain count,
 *  if equals 0, erase "this" pointer
 */
func ObjectRelease(obj interface{}) int {
	o, ok := obj.(SelfReference)
	if ok {
		// call 'Release()' from child class
		s := o.Self()
		if s == nil {
			return o.Release()
		} else {
			return s.Release()
		}
	} else {
		return 0
	}
}

/**
 *  Append the object to AutoreleasePool
 */
func ObjectAutorelease(obj interface{}) SelfReference {
	o, ok := obj.(SelfReference)
	if ok {
		// call 'Autorelease()' from child class
		s := o.Self()
		if s == nil {
			return o.Autorelease()
		} else {
			return s.Autorelease()
		}
	} else {
		return o
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
